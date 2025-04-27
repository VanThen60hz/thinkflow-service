package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) DeleteUserNotes(ctx context.Context, req *pb.DeleteUserNotesReq) (*pb.DeleteUserNotesResp, error) {
	deletedCount, err := s.business.DeleteUserNotes(ctx, int(req.UserId))
	if err != nil {
		return &pb.DeleteUserNotesResp{
			Success:      false,
			DeletedCount: 0,
		}, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.DeleteUserNotesResp{
		Success:      true,
		DeletedCount: int32(deletedCount),
	}, nil
}
