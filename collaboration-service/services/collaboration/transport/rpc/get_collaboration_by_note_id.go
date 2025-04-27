package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetCollaborationByNoteId(ctx context.Context, req *pb.GetCollaborationByNoteIdRequest) (*pb.GetCollaborationByNoteIdResponse, error) {
	paging := &core.Paging{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}

	collabs, err := s.business.GetCollaborationByNoteId(ctx, int(req.NoteId), paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	collaborations := make([]*pb.Collaboration, len(collabs))
	for i, collab := range collabs {
		collaborations[i] = toPBCollaboration(&collab)
	}

	return &pb.GetCollaborationByNoteIdResponse{
		Collaborations: collaborations,
		Total:          int32(len(collaborations)),
	}, nil
}
