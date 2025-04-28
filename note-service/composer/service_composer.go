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

	noteShareLinkSQLRepository "thinkflow-service/services/note-share-links/repository/mysql"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/gin-gonic/gin"
)

type NoteService interface {
	CreateNoteHdl() func(*gin.Context)
	CreateNoteShareLinkHdl() func(*gin.Context)
	NoteShareLinkToEmailHdl() func(*gin.Context)
	AcceptSharedNoteHdl() func(*gin.Context)
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
	GetTextHdl() func(*gin.Context)
	GetTextByNoteIdHdl() func(*gin.Context)
	UpdateTextHdl() func(*gin.Context)
	DeleteTextHdl() func(*gin.Context)
}

func ComposeNoteAPIService(serviceCtx sctx.ServiceContext) NoteService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	jwtProvider := serviceCtx.MustGet(common.KeyCompJWT).(common.JWTProvider)

	userClient := noteRepoRPC.NewClient(composeUserRPCClient(serviceCtx))

	summaryClient := noteRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))
	mindmapClient := noteRepoRPC.NewMindmapClient(ComposeMindmapRPCClient(serviceCtx))
	collabClient := noteRepoRPC.NewCollaborationClient(ComposeCollaborationRPCClient(serviceCtx))

	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	noteShareLinkRepo := noteShareLinkSQLRepository.NewMySQLRepository(db.GetDB())

	redisClient := serviceCtx.MustGet(common.KeyCompRedis).(redisc.Redis)
	emailService := serviceCtx.MustGet(common.KeyCompEmail).(emailc.Email)

	biz := noteBusiness.NewBusiness(
		noteRepo, userClient, collabClient,
		noteShareLinkRepo,
		summaryClient, mindmapClient,
		jwtProvider, redisClient, emailService,
	)
	serviceAPI := noteAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeTextAPIService(serviceCtx sctx.ServiceContext) TextService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	summaryClient := textRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))
	// collabClient := textRepoRPC.NewCollaborationClient(ComposeCollaborationRPCClient(serviceCtx))

	textRepo := textSQLRepository.NewMySQLRepository(db.GetDB())
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())

	// biz := textBusiness.NewBusiness(textRepo, noteRepo, collabClient, summaryClient)
	biz := textBusiness.NewBusiness(textRepo, noteRepo, summaryClient)
	serviceAPI := textAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeNoteGRPCService(serviceCtx sctx.ServiceContext) pb.NoteServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	noteBiz := noteBusiness.NewBusiness(noteRepo, nil, nil, nil, nil, nil, nil, nil, nil)
	return noteRPC.NewService(noteBiz)
}
