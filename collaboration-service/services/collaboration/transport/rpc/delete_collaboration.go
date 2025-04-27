package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) DeleteCollaboration(ctx context.Context, req *pb.DeleteCollaborationRequest) (*pb.DeleteCollaborationResponse, error) {
	if err := s.business.DeleteCollaboration(ctx, int(req.Id)); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.DeleteCollaborationResponse{Success: true}, nil
}
