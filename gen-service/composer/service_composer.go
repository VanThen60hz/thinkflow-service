package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"

	transcriptBusiness "thinkflow-service/services/transcript/business"
	transcriptSQLRepository "thinkflow-service/services/transcript/repository/mysql"
	transcriptAPI "thinkflow-service/services/transcript/transport/api"
	transcriptRPC "thinkflow-service/services/transcript/transport/rpc"

	summaryBusiness "thinkflow-service/services/summary/business"
	summarySQLRepository "thinkflow-service/services/summary/repository/mysql"
	summaryAPI "thinkflow-service/services/summary/transport/api"
	summaryRPC "thinkflow-service/services/summary/transport/rpc"

	mindmapBusiness "thinkflow-service/services/mindmap/business"
	mindmapSQLRepository "thinkflow-service/services/mindmap/repository/mysql"
	mindmapAPI "thinkflow-service/services/mindmap/transport/api"
	mindmapRPC "thinkflow-service/services/mindmap/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type GenService struct {
	Transcript TranscriptService
	Summary    SummaryService
	Mindmap    MindmapService
}

type TranscriptService interface {
	UpdateTranscriptHdl() func(*gin.Context)
	DeleteTranscriptHdl() func(*gin.Context)
	GetTranscriptHdl() func(*gin.Context)
}

type SummaryService interface {
	UpdateSummaryHdl() func(*gin.Context)
	DeleteSummaryHdl() func(*gin.Context)
	GetSummaryHdl() func(*gin.Context)
}

type MindmapService interface {
	UpdateMindmapHdl() func(*gin.Context)
	DeleteMindmapHdl() func(*gin.Context)
	GetMindmapHdl() func(*gin.Context)
}

func ComposeGenAPIService(serviceCtx sctx.ServiceContext) GenService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	transcriptRepo := transcriptSQLRepository.NewMySQLRepository(db.GetDB())
	transcriptBiz := transcriptBusiness.NewBusiness(transcriptRepo)
	transcriptService := transcriptAPI.NewAPI(serviceCtx, transcriptBiz)

	summaryRepo := summarySQLRepository.NewMySQLRepository(db.GetDB())
	summaryBiz := summaryBusiness.NewBusiness(summaryRepo)
	summaryService := summaryAPI.NewAPI(serviceCtx, summaryBiz)

	mindmapRepo := mindmapSQLRepository.NewMySQLRepository(db.GetDB())
	mindmapBiz := mindmapBusiness.NewBusiness(mindmapRepo)
	mindmapService := mindmapAPI.NewAPI(serviceCtx, mindmapBiz)

	return GenService{
		Transcript: transcriptService,
		Summary:    summaryService,
		Mindmap:    mindmapService,
	}
}

func ComposeTranscriptGRPCService(serviceCtx sctx.ServiceContext) pb.TranscriptServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	transcriptRepo := transcriptSQLRepository.NewMySQLRepository(db.GetDB())
	transcriptBiz := transcriptBusiness.NewBusiness(transcriptRepo)
	return transcriptRPC.NewService(transcriptBiz)
}

func ComposeSummaryGRPCService(serviceCtx sctx.ServiceContext) pb.SummaryServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	summaryRepo := summarySQLRepository.NewMySQLRepository(db.GetDB())
	summaryBiz := summaryBusiness.NewBusiness(summaryRepo)
	return summaryRPC.NewService(summaryBiz)
}

func ComposeMindmapGRPCService(serviceCtx sctx.ServiceContext) pb.MindmapServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	mindmapRepo := mindmapSQLRepository.NewMySQLRepository(db.GetDB())
	mindmapBiz := mindmapBusiness.NewBusiness(mindmapRepo)
	return mindmapRPC.NewService(mindmapBiz)
}
