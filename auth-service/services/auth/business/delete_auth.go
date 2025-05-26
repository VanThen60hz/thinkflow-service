package business

import (
	"context"

	"github.com/pkg/errors"
)

func (biz *business) DeleteAuth(ctx context.Context, userId int) error {
	auth, err := biz.repository.GetAuthByUserId(ctx, userId)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := biz.repository.DeleteAuth(ctx, auth.Email); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
