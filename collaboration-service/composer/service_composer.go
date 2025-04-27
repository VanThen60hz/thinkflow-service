package composer

import (
	"thinkflow-service/common"
	"thinkflow-service/proto/pb"
	collabBusiness "thinkflow-service/services/collaboration/business"
	collabSQLRepository "thinkflow-service/services/collaboration/repository/mysql"
	collabRPC "thinkflow-service/services/collaboration/transport/rpc"

	sctx "github.com/VanThen60hz/service-context"
)

func ComposeCollaborationGRPCService(serviceCtx sctx.ServiceContext) pb.CollaborationServiceServer {
	db := serviceCtx.MustGet(common.KeyCompMySQL).(common.GormComponent)

	collabRepo := collabSQLRepository.NewMySQLRepository(db.GetDB())
	collabBiz := collabBusiness.NewBusiness(collabRepo)
	collabService := collabRPC.NewService(collabBiz)

	return collabService
}
