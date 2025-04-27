package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) DeleteCollaboration(ctx context.Context, id int) error {
	if err := biz.collabRepo.DeleteCollaboration(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError("Cannot delete collaboration").
			WithDebug(err.Error())
	}

	return nil
} 