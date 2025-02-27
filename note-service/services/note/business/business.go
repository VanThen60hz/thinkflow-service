package business

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
)

type NoteRepository interface {
	AddNewNote(ctx context.Context, data *entity.NoteDataCreation) error
	UpdateNote(ctx context.Context, id int, data *entity.NoteDataUpdate) error
	DeleteNote(ctx context.Context, id int) error
	GetNoteById(ctx context.Context, id int) (*entity.Note, error)
	ListNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error)
}

type UserRepository interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]core.SimpleUser, error)
	GetUserById(ctx context.Context, id int) (*core.SimpleUser, error)
}

type business struct {
	noteRepo NoteRepository
	userRepo UserRepository
}

func NewBusiness(noteRepo NoteRepository, userRepo UserRepository) *business {
	return &business{
		noteRepo: noteRepo,
		userRepo: userRepo,
	}
}
