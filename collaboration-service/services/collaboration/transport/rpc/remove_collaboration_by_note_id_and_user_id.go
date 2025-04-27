package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) RemoveCollaborationByNoteIdAndUserId(ctx context.Context, req *pb.RemoveCollaborationByNoteIdAndUserIdRequest) (*pb.RemoveCollaborationByNoteIdAndUserIdResponse, error) {
	if err := s.business.RemoveCollaborationByNoteIdAndUserId(ctx, int(req.NoteId), int(req.UserId)); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.RemoveCollaborationByNoteIdAndUserIdResponse{Success: true}, nil
}
