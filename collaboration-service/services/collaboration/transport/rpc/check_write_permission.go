package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) CheckWritePermission(ctx context.Context, req *pb.CheckWritePermissionRequest) (*pb.CheckWritePermissionResponse, error) {
	hasPermission, err := s.business.HasWritePermission(ctx, int(req.NoteId), int(req.UserId))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.CheckWritePermissionResponse{HasPermission: hasPermission}, nil
}
