package rpc

import (
	"context"

	"thinkflow-service/services/note/entity"
)

type Business interface {
	DeleteUserNotes(ctx context.Context, userId int) (int, error)
	GetNoteByIdInt64(ctx context.Context, noteId int64) (*entity.Note, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
