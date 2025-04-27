package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) UpdateCollaboration(ctx context.Context, req *pb.UpdateCollaborationRequest) (*pb.UpdateCollaborationResponse, error) {
	collabData := &entity.Collaboration{
		Permission: toEntityPermissionType(req.Collaboration.Permission),
	}

	if err := s.business.UpdateCollaboration(ctx, int(req.Id), collabData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.UpdateCollaborationResponse{Success: true}, nil
}
