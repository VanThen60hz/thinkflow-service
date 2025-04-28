package business

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (biz *business) GetNoteShareLinkByID(ctx context.Context, id int64) (*entity.NoteShareLink, error) {
	link, err := biz.noteShareLinkRepo.GetNoteShareLinkByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return link, nil
}
