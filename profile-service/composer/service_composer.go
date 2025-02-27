package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	userBusiness "thinkflow-service/services/user/business"
	userSQLRepository "thinkflow-service/services/user/repository/mysql"
	userApi "thinkflow-service/services/user/transport/api"
	userRPC "thinkflow-service/services/user/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserProfileHdl() func(*gin.Context)
}

func ComposeUserAPIService(serviceCtx sctx.ServiceContext) UserService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	biz := userBusiness.NewBusiness(userRepo)
	userService := userApi.NewAPI(biz)

	return userService
}

func ComposeUserGRPCService(serviceCtx sctx.ServiceContext) pb.UserServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userRepo := userSQLRepository.NewMySQLRepository(db.GetDB())
	userBiz := userBusiness.NewBusiness(userRepo)
	userService := userRPC.NewService(userBiz)

	return userService
}
