package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error {
	err := biz.userRepo.CreateNewUser(ctx, data)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateUser.Error()).
			WithDebug(err.Error())
	}

	return nil
}
