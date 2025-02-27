package composer

import (
	"thinkflow-service/common"
	noteBusiness "thinkflow-service/services/note/business"
	noteSQLRepository "thinkflow-service/services/note/repository/mysql"
	noteUserRPC "thinkflow-service/services/note/repository/rpc"
	noteAPI "thinkflow-service/services/note/transport/api"

	sctx "github.com/VanThen60hz/service-context"
	"github.com/gin-gonic/gin"
)

type NoteService interface {
	CreateNoteHdl() func(*gin.Context)
	GetNoteHdl() func(*gin.Context)
	ListNoteHdl() func(*gin.Context)
	UpdateNoteHdl() func(*gin.Context)
	DeleteNoteHdl() func(*gin.Context)
	DoneNoteHdl() func(*gin.Context)
	DoingNoteHdl() func(*gin.Context)
}

func ComposeNoteAPIService(serviceCtx sctx.ServiceContext) NoteService {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	userClient := noteUserRPC.NewClient(composeUserRPCClient(serviceCtx))
	noteRepo := noteSQLRepository.NewMySQLRepository(db.GetDB())
	biz := noteBusiness.NewBusiness(noteRepo, userClient)
	serviceAPI := noteAPI.NewAPI(serviceCtx, biz)

	return serviceAPI
}
