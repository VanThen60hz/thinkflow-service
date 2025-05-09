package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	noteBusiness "thinkflow-service/services/note/business"
	noteSQLRepository "thinkflow-service/services/note/repository/mysql"
	noteRepoRPC "thinkflow-service/services/note/repository/rpc"
	noteAPI "thinkflow-service/services/note/transport/api"
	noteRPC "thinkflow-service/services/note/transport/rpc"

	textBusiness "thinkflow-service/services/text/business"
	textSQLRepository "thinkflow-service/services/text/repository/mysql"
	textRepoRPC "thinkflow-service/services/text/repository/rpc"
	textAPI "thinkflow-service/services/text/transport/api"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/component/s3c"
	"github.com/gin-gonic/gin"
)

type NoteService interface {
	CreateNoteHdl() func(*gin.Context)
	CreateNoteShareLinkHdl() func(*gin.Context)
	NoteShareLinkToEmailHdl() func(*gin.Context)
	AcceptSharedNoteHdl() func(*gin.Context)
	SummaryNoteHdl() func(*gin.Context)
	MindmapNoteHdl() func(*gin.Context)
	GetNoteHdl() func(*gin.Context)
	ListNotesHdl() func(*gin.Context)
	ListNoteMembersHdl() func(*gin.Context)
	ListNotesSharedWithMeHdl() func(*gin.Context)
	ListArchivedNotesHdl() func(*gin.Context)
	UpdateNoteHdl() func(*gin.Context)
	ArchiveNoteHdl() func(*gin.Context)
	UnarchiveNoteHdl() func(*gin.Context)
	UpdateNoteMemberHdl() func(*gin.Context)
	DeleteNoteHdl() func(*gin.Context)
	DeleteNoteMemberHdl() func(*gin.Context)
}

type TextService interface {
	CreateTextHdl() func(*gin.Context)
	SummaryTextHdl() func(*gin.Context)
	GetTextHdl() func(*gin.Context)
	GetTextByNoteIdHdl() func(*gin.Context)
	UpdateTextHdl() func(*gin.Context)
	DeleteTextHdl() func(*gin.Context)
}

func ComposeNoteAPIService(serviceCtx sctx.ServiceContext) NoteService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtProvider := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)
	s3Client := serviceCtx.MustGet(common.KeyCompS3).(*s3c.S3Component)

	userClient := noteRepoRPC.NewClient(composeUserRPCClient(serviceCtx))
	audioClient := noteRepoRPC.NewAudioClient(ComposeAudioRPCClient(serviceCtx))
	imageClient := noteRepoRPC.NewImageClient(ComposeImageRPCClient(serviceCtx))

	transcriptClient := noteRepoRPC.NewTranscriptClient(ComposeTranscriptRPCClient(serviceCtx))
	summaryClient := noteRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))
	mindmapClient := noteRepoRPC.NewMindmapClient(ComposeMindmapRPCClient(serviceCtx))
	collabClient := noteRepoRPC.NewCollaborationClient(ComposeCollaborationRPCClient(serviceCtx))
	noteShareLinkClient := noteRepoRPC.NewNoteShareLinkClient(ComposeNoteShareLinkRPCClient(serviceCtx))

	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	textRepo := textSQLRepository.NewMySQLRepository(db.GetDB())

	redisClient := serviceCtx.MustGet(common.KeyCompRedis).(redisc.Redis)
	emailService := serviceCtx.MustGet(common.KeyCompEmail).(emailc.Email)

	noteBiz := noteBusiness.NewBusiness(
		noteRepo, textRepo,
		userClient, imageClient, audioClient, collabClient, noteShareLinkClient,
		transcriptClient, summaryClient, mindmapClient,
		jwtProvider, s3Client, redisClient, emailService,
	)
	serviceAPI := noteAPI.NewAPI(serviceCtx, noteBiz)

	return serviceAPI
}

func ComposeTextAPIService(serviceCtx sctx.ServiceContext) TextService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	summaryClient := textRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))
	collabClient := textRepoRPC.NewCollaborationClient(ComposeCollaborationRPCClient(serviceCtx))

	textRepo := textSQLRepository.NewMySQLRepository(db.GetDB())
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())

	biz := textBusiness.NewBusiness(textRepo, noteRepo, collabClient, summaryClient)
	serviceAPI := textAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeNoteGRPCService(serviceCtx sctx.ServiceContext) pb.NoteServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	noteBiz := noteBusiness.NewBusiness(noteRepo, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	return noteRPC.NewService(noteBiz)
}
