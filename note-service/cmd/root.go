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
	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/ginc"
	smdlw "github.com/VanThen60hz/service-context/component/ginc/middleware"
	"github.com/VanThen60hz/service-context/component/gormc"
	"github.com/VanThen60hz/service-context/component/jwtc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("ThinkFlow Microservices"),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompMySQL, "")),
		sctx.WithComponent(jwtc.NewJWT(common.KeyCompJWT)),
		sctx.WithComponent(s3c.NewS3Component(common.KeyCompS3)),
		sctx.WithComponent(redisc.NewRedisComponent(common.KeyCompRedis)),
		sctx.WithComponent(emailc.NewEmailComponent(common.KeyCompEmail)),
		sctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
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
	noteAPIService := composer.ComposeNoteAPIService(serviceCtx)
	textAPIService := composer.ComposeTextAPIService(serviceCtx)

	requireAuthMdw := middleware.RequireAuth(composer.ComposeAuthRPCClient(serviceCtx))

	notes := router.Group("/notes", requireAuthMdw)
	{
		notes.POST("", noteAPIService.CreateNoteHdl())
		notes.POST("/:note-id/share/link", noteAPIService.CreateNoteShareLinkHdl())
		notes.POST(":note-id/share/email", noteAPIService.NoteShareLinkToEmailHdl())
		notes.POST("/accept/:token", noteAPIService.AcceptSharedNoteHdl())
		notes.POST("/:note-id/summary", noteAPIService.SummaryNoteHdl())
		notes.POST("/:note-id/mindmap", noteAPIService.MindmapNoteHdl())
		notes.GET("", noteAPIService.ListNotesHdl())
		notes.GET("/shared-with-me", noteAPIService.ListNotesSharedWithMeHdl())
		notes.GET("/archived", noteAPIService.ListArchivedNotesHdl())
		notes.GET("/:note-id", noteAPIService.GetNoteHdl())
		notes.GET("/:note-id/members", noteAPIService.ListNoteMembersHdl())
		notes.PATCH("/:note-id", noteAPIService.UpdateNoteHdl())
		notes.PATCH("/archive/:note-id", noteAPIService.ArchiveNoteHdl())
		notes.PATCH("/unarchive/:note-id", noteAPIService.UnarchiveNoteHdl())
		notes.PATCH("/:note-id/members/:user-id", noteAPIService.UpdateNoteMemberHdl())
		notes.DELETE("/:note-id", noteAPIService.DeleteNoteHdl())
		notes.DELETE("/:note-id/members/:user-id", noteAPIService.DeleteNoteMemberHdl())
	}

	texts := router.Group("/texts", requireAuthMdw)
	{
		texts.POST("/note/:note-id", textAPIService.CreateTextHdl())
		texts.POST("/:text-id/summary", textAPIService.SummaryTextHdl())
		texts.GET("/note/:note-id", textAPIService.GetTextByNoteIdHdl())
		texts.GET("/:text-id", textAPIService.GetTextHdl())
		texts.PATCH("/:text-id", textAPIService.UpdateTextHdl())
		texts.DELETE("/:text-id", textAPIService.DeleteTextHdl())
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

	pb.RegisterNoteServiceServer(s, composer.ComposeNoteGRPCService(serviceCtx))

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
