package composer

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/middleware"
	"thinkflow-service/proto/pb"
	userBusiness "thinkflow-service/services/user/business"
	userSQLRepository "thinkflow-service/services/user/repository/mysql"
	userRPCRepository "thinkflow-service/services/user/repository/rpc"
	userApi "thinkflow-service/services/user/transport/api"
	userRPC "thinkflow-service/services/user/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUserHdl() func(*gin.Context)
	DeactivateUserHdl() func(*gin.Context)
	GetDashboardStatsHdl() func(*gin.Context)
	GetUserProfileHdl() func(*gin.Context)
	ListUserHdl() func(*gin.Context)
	UpdateUserProfileHdl() func(*gin.Context)
	UpdateUserHdl() func(*gin.Context)
	DeleteUserHdl() func(*gin.Context)
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

func ComposeUserAPIService(serviceCtx sctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageClient := userRPCRepository.NewImageClient(ComposeImageRPCClient(serviceCtx))
	noteClient := userRPCRepository.NewNoteClient(ComposeNoteRPCClient(serviceCtx))
	authClient := userRPCRepository.NewAuthClient(ComposeAuthRPCClient(serviceCtx))

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())

	biz := userBusiness.NewBusiness(userRepo, imageClient, noteClient, authClient)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageClient := userRPCRepository.NewImageClient(ComposeImageRPCClient(serviceCtx))
	noteClient := userRPCRepository.NewNoteClient(ComposeNoteRPCClient(serviceCtx))
	authClient := userRPCRepository.NewAuthClient(ComposeAuthRPCClient(serviceCtx))

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())

	userBiz := userBusiness.NewBusiness(userRepo, imageClient, noteClient, authClient)
	userService := userRPC.NewService(userBiz)

	return userService
}
