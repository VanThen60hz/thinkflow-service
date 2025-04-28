package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	userBusiness "thinkflow-service/services/user/business"
	userSQLRepository "thinkflow-service/services/user/repository/mysql"
	userImageNoteRPC "thinkflow-service/services/user/repository/rpc"
	userApi "thinkflow-service/services/user/transport/api"
	userRPC "thinkflow-service/services/user/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserProfileHdl() func(*gin.Context)
	UpdateUserProfileHdl() func(*gin.Context)
	DeleteUserHdl() func(*gin.Context)
}

func ComposeUserAPIService(serviceCtx sctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageClient := userImageNoteRPC.NewImageClient(ComposeImageRPCClient(serviceCtx))
	noteClient := userImageNoteRPC.NewNoteClient(ComposeNoteRPCClient(serviceCtx))

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())

	biz := userBusiness.NewBusiness(userRepo, imageClient, noteClient)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageClient := userImageNoteRPC.NewImageClient(ComposeImageRPCClient(serviceCtx))
	noteClient := userImageNoteRPC.NewNoteClient(ComposeNoteRPCClient(serviceCtx))

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())

	userBiz := userBusiness.NewBusiness(userRepo, imageClient, noteClient)
	userService := userRPC.NewService(userBiz)

	return userService
}
