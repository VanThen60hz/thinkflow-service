package composer

import (
	"os"

	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	authBusiness "thinkflow-service/services/auth/business"
	authSQLRepository "thinkflow-service/services/auth/repository/mysql"
	authUserRPC "thinkflow-service/services/auth/repository/rpc"
	authAPI "thinkflow-service/services/auth/transport/api"
	authRPC "thinkflow-service/services/auth/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginHdl() func(*gin.Context)
	RegisterHdl() func(*gin.Context)
	ForgotPasswordHdl() func(*gin.Context)
	ResetPasswordHdl() func(*gin.Context)
	GoogleLoginHdl() func(*gin.Context)
	GoogleCallbackHdl() func(*gin.Context)
	FacebookLoginHdl() func(*gin.Context)
	FacebookCallbackHdl() func(*gin.Context)
	LogoutHdl() func(*gin.Context)
}

func ComposeAuthAPIService(serviceCtx sctx.ServiceContext) AuthService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)

	userClient := authUserRPC.NewClient(ComposeUserRPCClient(serviceCtx))
	userClient.SetDB(db.GetDB())
	redisClient := common.NewRedisClient(os.Getenv("REDIS_ADDRESS"))
	emailService := common.NewEmailService(os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	biz := authBusiness.NewBusiness(authRepo, userClient, jwtComp, hasher, redisClient, emailService)
	serviceAPI := authAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeAuthGRPCService(serviceCtx sctx.ServiceContext) pb.AuthServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtComp := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	authRepo := authSQLRepository.NewMySQLRepository(db.GetDB())
	hasher := new(common.Hasher)
	redisClient := common.NewRedisClient(os.Getenv("REDIS_ADDRESS"))
	emailService := common.NewEmailService(os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	biz := authBusiness.NewBusiness(authRepo, nil, jwtComp, hasher, redisClient, emailService)
	authService := authRPC.NewService(biz)

	return authService
}
