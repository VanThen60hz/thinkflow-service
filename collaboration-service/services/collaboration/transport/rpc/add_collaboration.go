package rpc

import (
	"context"

	"thinkflow-service/proto/pb"
	"thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) AddCollaboration(ctx context.Context, req *pb.AddCollaborationRequest) (*pb.AddCollaborationResponse, error) {
	collabData := &entity.CollaborationCreation{
		NoteId:     int(req.Collaboration.NoteId),
		UserId:     int(req.Collaboration.UserId),
		Permission: toEntityPermissionType(req.Collaboration.Permission),
	}

	if err := s.business.AddNewCollaboration(ctx, collabData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.AddCollaborationResponse{Success: true}, nil
}
