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
	"thinkflow-service/services/notification/repository"
	"thinkflow-service/services/notification/transport/fcm"
	"thinkflow-service/services/notification/transport/ws"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/ginc"
	smdlw "github.com/VanThen60hz/service-context/component/ginc/middleware"
	"github.com/VanThen60hz/service-context/component/gormc"
	"github.com/VanThen60hz/service-context/component/jwtc"
	"github.com/VanThen60hz/service-context/component/natsc"
	"github.com/VanThen60hz/service-context/core"
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
		sctx.WithComponent(natsc.NewNatsComponent(common.KeyCompNats)),
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
		natsComp := serviceCtx.MustGet(common.KeyCompNats).(natsc.Nats)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx))

		router.Use(middleware.Cors())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent).GetDB()
		fcmTokenRepo := repository.NewFCMTokenRepository(db)

		// Initialize FCM service
		fcmService, err := fcm.NewService("config/firebase-credentials.json", fcmTokenRepo)
		if err != nil {
			logger.Fatal(err)
		}

		// Add FCM service to context
		router.Use(func(c *gin.Context) {
			c.Set("fcm_service", fcmService)
			c.Next()
		})

		// Start WebSocket hub
		go ws.Hub.Run()

		// Start NATS subscriber with FCM service
		go ws.StartNatsSubscriber(cmd.Context(), natsComp, fcmService)

		go StartGRPCServices(serviceCtx)

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx sctx.ServiceContext) {
	notiAPIService := composer.ComposeNotificationAPIService(serviceCtx)

	requireAuthMdw := middleware.RequireAuth(composer.ComposeAuthClientForMiddleware(serviceCtx))

	notis := router.Group("/notifications", requireAuthMdw)
	{
		notis.POST("", notiAPIService.CreateNotification)
		notis.GET("/unread-count", notiAPIService.GetUnreadCountHdl())
		notis.GET("", notiAPIService.ListNotificationsHdl())
		notis.PATCH("/:noti-id/read", notiAPIService.MarkNotificationAsReadHdl())
		notis.PATCH("/read-all", notiAPIService.MarkAllNotificationsAsReadHdl())
		notis.DELETE("/:noti-id", notiAPIService.DeleteNotificationHdl())
		notis.GET("/ws", func(c *gin.Context) {
			ws.HandleWebSocket(c.Writer, c.Request)
		})
		notis.POST("/fcm-token", func(c *gin.Context) {
			var req struct {
				Token    string `json:"token" binding:"required"`
				DeviceID string `json:"device_id" binding:"required"`
				Platform string `json:"platform" binding:"required"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			requester := c.MustGet(common.RequesterKey).(core.Requester)
			userID := requester.GetSubject()

			fcmService := c.MustGet("fcm_service").(*fcm.Service)
			err := fcmService.RegisterToken(c.Request.Context(), userID, req.Token, req.DeviceID, req.Platform)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Token registered successfully"})
		})

		// Add new endpoint for unregistering FCM token
		notis.DELETE("/fcm-token", func(c *gin.Context) {
			var req struct {
				Token string `json:"token" binding:"required"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			fcmService := c.MustGet("fcm_service").(*fcm.Service)
			err := fcmService.UnregisterToken(c.Request.Context(), req.Token)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Token unregistered successfully"})
		})
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

	pb.RegisterNotificationServiceServer(s, composer.ComposeNotiGRPCService(serviceCtx))

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
