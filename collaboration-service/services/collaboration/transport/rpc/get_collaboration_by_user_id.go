package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetCollaborationByUserId(ctx context.Context, req *pb.GetCollaborationByUserIdRequest) (*pb.GetCollaborationByUserIdResponse, error) {
	paging := &core.Paging{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}

	collabs, err := s.business.GetCollaborationByUserId(ctx, int(req.UserId), paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	collaborations := make([]*pb.Collaboration, len(collabs))
	for i, collab := range collabs {
		collaborations[i] = toPBCollaboration(&collab)
	}

	return &pb.GetCollaborationByUserIdResponse{
		Collaborations: collaborations,
		Total:          int32(paging.Total),
	}, nil
}
