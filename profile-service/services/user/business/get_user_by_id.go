package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error) {
	users, err := biz.userRepo.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, core.ErrNotFound.
			WithError(entity.ErrCannotGetUsers.Error()).
			WithDebug(err.Error())
	}

	return users, nil
}
