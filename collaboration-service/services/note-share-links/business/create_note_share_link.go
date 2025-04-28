package business

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (biz *business) CreateNoteShareLink(ctx context.Context, data *entity.NoteShareLinkCreation) (*entity.NoteShareLink, error) {
	if err := biz.noteShareLinkRepo.AddNewNoteShareLink(ctx, data); err != nil {
		return nil, errors.WithStack(err)
	}

	link, err := biz.noteShareLinkRepo.GetNoteShareLinkByToken(ctx, data.Token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return link, nil
}
