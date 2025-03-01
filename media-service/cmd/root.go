package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/composer"
	"thinkflow-service/middleware"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/ginc"
	smdlw "github.com/VanThen60hz/service-context/component/ginc/middleware"
	"github.com/VanThen60hz/service-context/component/gormc"
	"github.com/VanThen60hz/service-context/component/jwtc"
	"github.com/VanThen60hz/service-context/component/s3c"

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

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	mediaAPIService := composer.ComposeMediaAPIService(serviceCtx)

	requireAuthMdw := middleware.RequireAuth(composer.ComposeAuthRPCClient(serviceCtx))

	media := router.Group("/media", requireAuthMdw)
	{
		images := media.Group("/images")
		{
			images.GET("", mediaAPIService.ListImagesHdl())
			images.POST("", mediaAPIService.CreateImageHdl())
			images.GET("/:image-id", mediaAPIService.GetImageHdl())
			images.PATCH("/:image-id", mediaAPIService.UpdateImageHdl())
			images.DELETE("/:image-id", mediaAPIService.DeleteImageHdl())
		}

		audios := media.Group("/audios")
		{
			audios.GET("", mediaAPIService.ListAudiosHdl())
			audios.POST("", mediaAPIService.CreateAudioHdl())
			audios.GET("/:audio-id", mediaAPIService.GetAudioHdl())
			audios.PATCH("/:audio-id", mediaAPIService.UpdateAudioHdl())
			audios.DELETE("/:audio-id", mediaAPIService.DeleteAudioHdl())
		}
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
