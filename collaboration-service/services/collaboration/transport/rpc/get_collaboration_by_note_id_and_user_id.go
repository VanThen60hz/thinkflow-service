package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetCollaborationByNoteIdAndUserId(ctx context.Context, req *pb.GetCollaborationByNoteIdAndUserIdRequest) (*pb.GetCollaborationByNoteIdAndUserIdResponse, error) {
	collab, err := s.business.GetCollaborationByNoteIdAndUserId(ctx, int(req.NoteId), int(req.UserId))
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.WithError("Collaboration not found")
		}
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.GetCollaborationByNoteIdAndUserIdResponse{
		Collaboration: toPBCollaboration(collab),
	}, nil
}
