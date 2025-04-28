package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	collabBusiness "thinkflow-service/services/collaboration/business"
	collabSQLRepository "thinkflow-service/services/collaboration/repository/mysql"
	collabRPC "thinkflow-service/services/collaboration/transport/rpc"

	noteShareLinkBusiness "thinkflow-service/services/note-share-links/business"
	noteShareLinkSQLRepository "thinkflow-service/services/note-share-links/repository/mysql"
	noteShareLinkRPC "thinkflow-service/services/note-share-links/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
)

func ComposeCollaborationGRPCService(serviceCtx sctx.ServiceContext) pb.CollaborationServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	collabRepo := collabSQLRepository.NewMySQLRepository(db.GetDB())
	collabBiz := collabBusiness.NewBusiness(collabRepo)
	collabService := collabRPC.NewService(collabBiz)

	return collabService
}

func ComposeNoteShareLinkGRPCService(serviceCtx sctx.ServiceContext) pb.NoteShareLinkServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	noteShareLinkRepo := noteShareLinkSQLRepository.NewMySQLRepository(db.GetDB())
	noteShareLinkBiz := noteShareLinkBusiness.NewBusiness(noteShareLinkRepo)
	noteShareLinkService := noteShareLinkRPC.NewService(noteShareLinkBiz)

	return noteShareLinkService
}
