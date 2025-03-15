package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	audioBusiness "thinkflow-service/services/audio/business"
	audioSQLRepository "thinkflow-service/services/audio/repository/mysql"
	audioAPI "thinkflow-service/services/audio/transport/api"

	imageBusiness "thinkflow-service/services/image/business"
	imageSQLRepository "thinkflow-service/services/image/repository/mysql"
	imageAPI "thinkflow-service/services/image/transport/api"
	imageRPC "thinkflow-service/services/image/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type MediaService struct {
	Image ImageService
	Audio AudioService
}

type ImageService interface {
	CreateImageHdl() func(*gin.Context)
	GetImageHdl() func(*gin.Context)
	ListImagesHdl() func(*gin.Context)
	UpdateImageHdl() func(*gin.Context)
	DeleteImageHdl() func(*gin.Context)
}

type AudioService interface {
	CreateAudioHdl() func(*gin.Context)
	GetAudioHdl() func(*gin.Context)
	ListAudiosHdl() func(*gin.Context)
	UpdateAudioHdl() func(*gin.Context)
	DeleteAudioHdl() func(*gin.Context)
}

func ComposeMediaAPIService(serviceCtx sctx.ServiceContext) MediaService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageRepo := imageSQLRepository.NewMySQLRepository(db.GetDB())
	imageBiz := imageBusiness.NewBusiness(imageRepo)
	imageService := imageAPI.NewAPI(serviceCtx, imageBiz)

	audioRepo := audioSQLRepository.NewMySQLRepository(db.GetDB())
	audioBiz := audioBusiness.NewBusiness(audioRepo)
	audioService := audioAPI.NewAPI(serviceCtx, audioBiz)

	return MediaService{
		Image: imageService,
		Audio: audioService,
	}
}

func ComposeMediaGRPCService(serviceCtx sctx.ServiceContext) pb.MediaServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	imageRepo := imageSQLRepository.NewMySQLRepository(db.GetDB())
	imageBiz := imageBusiness.NewBusiness(imageRepo)
	imageService := imageRPC.NewService(imageBiz)

	return imageService
}
