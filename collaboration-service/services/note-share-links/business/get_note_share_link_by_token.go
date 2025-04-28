package business

import (
	"context"

	"thinkflow-service/services/note-share-links/entity"

	"github.com/pkg/errors"
)

func (biz *business) GetNoteShareLinkByToken(ctx context.Context, token string) (*entity.NoteShareLink, error) {
	link, err := biz.noteShareLinkRepo.GetNoteShareLinkByToken(ctx, token)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return link, nil
}
