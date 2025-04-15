package business

import (
	"context"
	"errors"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	userId, err := biz.userRepo.GetUserIdByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return 0, core.ErrNotFound
		}
		return 0, core.ErrInternalServerError.WithError(entity.ErrCannotGetUser.Error()).WithDebug(err.Error())
	}

	return userId, nil
}
