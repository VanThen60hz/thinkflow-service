package business

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (biz *business) UpdateNoteShareLink(ctx context.Context, id int64, data *entity.NoteShareLinkUpdate) (*entity.NoteShareLink, error) {
	if err := biz.noteShareLinkRepo.UpdateNoteShareLink(ctx, id, data); err != nil {
		return nil, errors.WithStack(err)
	}

	link, err := biz.noteShareLinkRepo.GetNoteShareLinkByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return link, nil
}
