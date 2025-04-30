package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/composer"
	"thinkflow-service/middleware"
	"thinkflow-service/proto/pb"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/ginc"
	smdlw "github.com/VanThen60hz/service-context/component/ginc/middleware"
	"github.com/VanThen60hz/service-context/component/gormc"
	"github.com/VanThen60hz/service-context/component/jwtc"
	"github.com/VanThen60hz/service-context/component/s3c"
	"google.golang.org/grpc"

	// "github.com/VanThen60hz/service-context/component/s3c"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("ThinkFlow Media Service"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompMySQL, "")),
		sctx.WithComponent(jwtc.NewJWT(common.KeyCompJWT)),
		sctx.WithComponent(s3c.NewS3Component(common.KeyCompS3)),
		sctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start media service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		// Make some delay for DB ready (migration)
		// remove it if you already had your own DB
		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(common.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx))

		router.Use(middleware.Cors())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		go StartGRPCServices(serviceCtx)

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	imageAPIService := composer.ComposeImageAPIService(serviceCtx)
	audioAPIService := composer.ComposeAudioAPIService(serviceCtx)
	attachmentAPIService := composer.ComposeAttachmentAPIService(serviceCtx)

	requireAuthMdw := middleware.RequireAuth(composer.ComposeAuthRPCClient(serviceCtx))

	media := router.Group("/media", requireAuthMdw)
	{
		images := media.Group("/images")
		{
			images.POST("", imageAPIService.UploadImageHdl())
			images.GET("", imageAPIService.ListImagesHdl())
			images.GET("/:image-id", imageAPIService.GetImageHdl())
			images.PATCH("/:image-id", imageAPIService.UpdateImageHdl())
			images.DELETE("/:image-id", imageAPIService.DeleteImageHdl())
		}

		audios := media.Group("/audios")
		{
			audios.POST("/:note-id", audioAPIService.UploadAudioHdl())
			audios.GET("", audioAPIService.ListAudiosHdl())
			audios.GET("/notes/:note-id", audioAPIService.GetAudiosByNoteHdl())
			audios.GET("/:audio-id", audioAPIService.GetAudioHdl())
			audios.PATCH("/:audio-id", audioAPIService.UpdateAudioHdl())
			audios.DELETE("/:audio-id", audioAPIService.DeleteAudioHdl())
		}

		attachments := media.Group("/attachments")
		{
			attachments.POST("", attachmentAPIService.UploadAttachmentHdl())
			attachments.GET("/:attachment-id", attachmentAPIService.GetAttachmentHdl())
			attachments.GET("/notes/:note-id", attachmentAPIService.GetAttachmentsByNoteIDHdl())
			attachments.DELETE("/:attachment-id", attachmentAPIService.DeleteAttachmentHdl())
			attachments.PATCH("/:attachment-id", attachmentAPIService.UpdateAttachmentHdl())
		}
	}
}

func StartGRPCServices(serviceCtx sctx.ServiceContext) {
	configComp := serviceCtx.MustGet(common.KeyCompConf).(common.Config)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))
	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterImageServiceServer(s, composer.ComposeImageGRPCService(serviceCtx))
	pb.RegisterAudioServiceServer(s, composer.ComposeAudioGRPCService(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
