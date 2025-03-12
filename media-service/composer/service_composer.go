package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	mediaBusiness "thinkflow-service/services/media/business"
	mediaSQLRepository "thinkflow-service/services/media/repository/mysql"
	mediaAPI "thinkflow-service/services/media/transport/api"
	mediaRPC "thinkflow-service/services/media/transport/rpc"

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

	mediaRepo := mediaSQLRepository.NewMySQLRepository(db.GetDB())
	biz := mediaBusiness.NewBusiness(mediaRepo)
	serviceAPI := mediaAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeMediaGRPCService(serviceCtx sctx.ServiceContext) pb.ImageServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	mediaRepo := mediaSQLRepository.NewMySQLRepository(db.GetDB())
	mediaBiz := mediaBusiness.NewBusiness(mediaRepo)
	mediaService := mediaRPC.NewService(mediaBiz)

	return mediaService
}
