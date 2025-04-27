package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) CheckReadPermission(ctx context.Context, req *pb.CheckReadPermissionRequest) (*pb.CheckReadPermissionResponse, error) {
	hasPermission, err := s.business.HasReadPermission(ctx, int(req.NoteId), int(req.UserId))
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.CheckReadPermissionResponse{HasPermission: hasPermission}, nil
}
