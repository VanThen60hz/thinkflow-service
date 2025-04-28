package rpc

import (
	"context"

	"thinkflow-service/proto/pb"

	"github.com/VanThen60hz/service-context/core"
)

func (s *grpcService) GetNoteById(ctx context.Context, req *pb.GetNoteByIdReq) (*pb.GetNoteByIdResp, error) {
	noteId, err := core.FromBase58(req.NoteId)
	if err != nil {
		return nil, core.ErrBadRequest.WithError(err.Error())
	}

	note, err := s.business.GetNoteById(ctx, int(noteId.GetLocalID()))
	if err != nil {
		return nil, err
	}

	var summaryId, mindmapId int64
	if note.SummaryID != nil {
		summaryId = *note.SummaryID
	}
	if note.MindmapID != nil {
		mindmapId = *note.MindmapID
	}

	return &pb.GetNoteByIdResp{
		Id:         int64(note.Id),
		Title:      note.Title,
		Archived:   note.Archived,
		UserId:     int64(note.UserId),
		Permission: note.Permission,
		SummaryId:  summaryId,
		MindmapId:  mindmapId,
		CreatedAt:  note.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  note.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
