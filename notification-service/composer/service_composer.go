package composer

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/middleware"
	"thinkflow-service/proto/pb"
	notiBusiness "thinkflow-service/services/notification/business"
	notiSQLRepository "thinkflow-service/services/notification/repository/mysql"
	notiRPCRepository "thinkflow-service/services/notification/repository/rpc"

	notiApi "thinkflow-service/services/notification/transport/api"
	notiRPC "thinkflow-service/services/notification/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type NotificationService interface {
	GetUnreadCountHdl() func(*gin.Context)
	ListNotificationsHdl() func(*gin.Context)
	MarkNotificationAsReadHdl() func(*gin.Context)
	MarkAllNotificationsAsReadHdl() func(*gin.Context)
	DeleteNotificationHdl() func(*gin.Context)
}

type authClientWrapper struct {
	client pb.AuthServiceClient
}

func (w *authClientWrapper) IntrospectToken(ctx context.Context, accessToken string) (string, string, error) {
	resp, err := w.client.IntrospectToken(ctx, &pb.IntrospectReq{AccessToken: accessToken})
	if err != nil {
		return "", "", err
	}
	return resp.Sub, resp.Tid, nil
}

func ComposeAuthClientForMiddleware(serviceCtx sctx.ServiceContext) middleware.AuthClient {
	return &authClientWrapper{client: ComposeAuthRPCClient(serviceCtx)}
}

func ComposeNotificationAPIService(serviceCtx sctx.ServiceContext) NotificationService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	authClient := notiRPCRepository.NewAuthClient(ComposeAuthRPCClient(serviceCtx))

	notiRepo := notiSQLRepository.NewMySQLRepository(db.GetDB())

	biz := notiBusiness.NewBusiness(notiRepo, authClient)
	notiService := notiApi.NewAPI(biz)

	return notiService
}

func ComposeNotiGRPCService(serviceCtx sctx.ServiceContext) pb.NotificationServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	authClient := notiRPCRepository.NewAuthClient(ComposeAuthRPCClient(serviceCtx))

	notiRepo := notiSQLRepository.NewMySQLRepository(db.GetDB())

	notiBiz := notiBusiness.NewBusiness(notiRepo, authClient)
	notiService := notiRPC.NewService(notiBiz)

	return notiService
}
