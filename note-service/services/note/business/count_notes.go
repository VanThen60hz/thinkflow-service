package business

import (
	"context"
)

func (biz *business) CountNotes(ctx context.Context) (int64, error) {
	return biz.noteRepo.CountNotes(ctx)
} 