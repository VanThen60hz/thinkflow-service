package business

import (
	"context"

	"thinkflow-service/services/note/entity"
)

func (b *business) GetNoteByIdInt64(ctx context.Context, id int64) (*entity.Note, error) {
	return b.noteRepo.GetNoteById(ctx, int(id))
}
