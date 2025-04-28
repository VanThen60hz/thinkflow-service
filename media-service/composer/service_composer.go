package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	imageBusiness "thinkflow-service/services/image/business"
	imageSQLRepository "thinkflow-service/services/image/repository/mysql"
	imageAPI "thinkflow-service/services/image/transport/api"
	imageRPC "thinkflow-service/services/image/transport/rpc"

	audioBusiness "thinkflow-service/services/audio/business"
	audioSQLRepository "thinkflow-service/services/audio/repository/mysql"
	audioAPI "thinkflow-service/services/audio/transport/api"
	audioRPC "thinkflow-service/services/audio/transport/rpc"

	audioRepoRPC "thinkflow-service/services/audio/repository/rpc"

	attachmentBusiness "thinkflow-service/services/attachment/business"
	attachmentSQLRepository "thinkflow-service/services/attachment/repository/mysql"
	attachmentRepoRPC "thinkflow-service/services/attachment/repository/rpc"
	attachmentAPI "thinkflow-service/services/attachment/transport/api"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/gin-gonic/gin"
)

type MediaService struct {
	Image      ImageService
	Audio      AudioService
	Attachment AttachmentService
}

type ImageService interface {
	UploadImageHdl() func(*gin.Context)
	GetImageHdl() func(*gin.Context)
	ListImagesHdl() func(*gin.Context)
	UpdateImageHdl() func(*gin.Context)
	DeleteImageHdl() func(*gin.Context)
}

type AudioService interface {
	UploadAudioHdl() func(*gin.Context)
	GetAudioHdl() func(*gin.Context)
	ListAudiosHdl() func(*gin.Context)
	UpdateAudioHdl() func(*gin.Context)
	DeleteAudioHdl() func(*gin.Context)
	GetAudiosByNoteHdl() func(*gin.Context)
}

type AttachmentService interface {
	UploadAttachmentHdl() func(*gin.Context)
	GetAttachmentHdl() func(*gin.Context)
	GetAttachmentsByNoteIDHdl() func(*gin.Context)
	DeleteAttachmentHdl() func(*gin.Context)
	UpdateAttachmentHdl() func(*gin.Context)
}

func ComposeMediaAPIService(serviceCtx sctx.ServiceContext) MediaService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	s3Client := serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)
	noteClient := attachmentRepoRPC.NewNoteClient(ComposeNoteRPCClient(serviceCtx))

	imageRepo := imageSQLRepository.NewMySQLRepository(db.GetDB())
	imageBiz := imageBusiness.NewBusiness(imageRepo, s3Client)
	imageService := imageAPI.NewAPI(serviceCtx, imageBiz)

	transcriptClient := audioRepoRPC.NewTranscriptClient(ComposeTranscriptRPCClient(serviceCtx))
	summaryClient := audioRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))

	audioRepo := audioSQLRepository.NewMySQLRepository(db.GetDB())
	audioBiz := audioBusiness.NewBusiness(audioRepo, s3Client, transcriptClient, summaryClient)
	audioService := audioAPI.NewAPI(serviceCtx, audioBiz)

	attachmentRepo := attachmentSQLRepository.NewMySQLRepository(db.GetDB())
	attachmentBiz := attachmentBusiness.NewBusiness(attachmentRepo, noteClient, s3Client)
	attachmentService := attachmentAPI.NewAPI(serviceCtx, attachmentBiz)

	return MediaService{
		Image:      imageService,
		Audio:      audioService,
		Attachment: attachmentService,
	}
}

func ComposeImageGRPCService(serviceCtx sctx.ServiceContext) pb.ImageServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	s3Client := serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)

	imageRepo := imageSQLRepository.NewMySQLRepository(db.GetDB())
	imageBiz := imageBusiness.NewBusiness(imageRepo, s3Client)
	return imageRPC.NewService(imageBiz)
}

func ComposeAudioGRPCService(serviceCtx sctx.ServiceContext) pb.AudioServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	audioRepo := audioSQLRepository.NewMySQLRepository(db.GetDB())
	s3Client := serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)

	transcriptClient := audioRepoRPC.NewTranscriptClient(ComposeTranscriptRPCClient(serviceCtx))
	summaryClient := audioRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))

	audioBiz := audioBusiness.NewBusiness(audioRepo, s3Client, transcriptClient, summaryClient)
	return audioRPC.NewService(audioBiz)
}
