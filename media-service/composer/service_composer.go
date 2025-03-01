package composer

import (
	"thinkflow-service/common"
	mediaBusiness "thinkflow-service/services/media/business"
	mediaSQLRepository "thinkflow-service/services/media/repository/mysql"
	mediaUserRPC "thinkflow-service/services/media/repository/rpc"
	mediaAPI "thinkflow-service/services/media/transport/api"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type MediaService interface {
	CreateImageHdl() func(*gin.Context)
	CreateAudioHdl() func(*gin.Context)
	GetImageHdl() func(*gin.Context)
	GetAudioHdl() func(*gin.Context)
	ListImagesHdl() func(*gin.Context)
	ListAudiosHdl() func(*gin.Context)
	UpdateImageHdl() func(*gin.Context)
	UpdateAudioHdl() func(*gin.Context)
	DeleteImageHdl() func(*gin.Context)
	DeleteAudioHdl() func(*gin.Context)
}

func ComposeMediaAPIService(serviceCtx sctx.ServiceContext) MediaService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	userClient := ComposeUserRPCClient(serviceCtx)

	mediaRepo := mediaSQLRepository.NewMySQLRepository(db.GetDB())
	userRepo := mediaUserRPC.NewClient(userClient)
	biz := mediaBusiness.NewBusiness(mediaRepo, userRepo)
	serviceAPI := mediaAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}
