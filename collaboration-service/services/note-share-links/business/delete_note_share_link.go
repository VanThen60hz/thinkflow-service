package business

import (
	"context"

	"github.com/pkg/errors"
)

func (biz *business) DeleteNoteShareLink(ctx context.Context, id int64) error {
	if err := biz.noteShareLinkRepo.DeleteNoteShareLink(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
