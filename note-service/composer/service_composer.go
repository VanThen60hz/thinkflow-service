package composer

import (
	"thinkflow-service/common"
	noteBusiness "thinkflow-service/services/note/business"
	noteSQLRepository "thinkflow-service/services/note/repository/mysql"
	noteUserRPC "thinkflow-service/services/note/repository/rpc"
	noteAPI "thinkflow-service/services/note/transport/api"

	textBusiness "thinkflow-service/services/text/business"
	textSQLRepository "thinkflow-service/services/text/repository/mysql"
	textRepoRPC "thinkflow-service/services/text/repository/rpc"
	textAPI "thinkflow-service/services/text/transport/api"

	collabSQLRepository "thinkflow-service/services/collaboration/repository/mysql"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type NoteService interface {
	CreateNoteHdl() func(*gin.Context)
	GetNoteHdl() func(*gin.Context)
	ListNotesHdl() func(*gin.Context)
	ListNoteMembersHdl() func(*gin.Context)
	ListNotesSharedWithMeHdl() func(*gin.Context)
	ListArchivedNotesHdl() func(*gin.Context)
	UpdateNoteHdl() func(*gin.Context)
	ArchiveNoteHdl() func(*gin.Context)
	UnarchiveNoteHdl() func(*gin.Context)
	DeleteNoteHdl() func(*gin.Context)
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

	userClient := noteUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	collabRepo := collabSQLRepository.NewMySQLRepository(db.GetDB())

	biz := noteBusiness.NewBusiness(noteRepo, userClient, collabRepo)
	serviceAPI := noteAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}

func ComposeTextAPIService(serviceCtx sctx.ServiceContext) TextService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	summaryClient := textRepoRPC.NewSummaryClient(ComposeSummaryRPCClient(serviceCtx))
	mindmapClient := textRepoRPC.NewMindmapClient(ComposeMindmapRPCClient(serviceCtx))

	textRepo := textSQLRepository.NewMySQLRepository(db.GetDB())
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	collabRepo := collabSQLRepository.NewMySQLRepository(db.GetDB())

	biz := textBusiness.NewBusiness(textRepo, noteRepo, collabRepo, summaryClient, mindmapClient)
	serviceAPI := textAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}
