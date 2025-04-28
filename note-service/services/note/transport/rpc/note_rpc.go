package rpc

import (
	"context"

	"thinkflow-service/services/note/entity"
)

type Business interface {
	DeleteUserNotes(ctx context.Context, userId int) (int, error)
	GetNoteById(ctx context.Context, noteId int) (*entity.Note, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{business: business}
}
