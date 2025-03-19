package business

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetUserStatus(ctx context.Context, id int) (string, error) {
	user, err := biz.userRepo.GetUserById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return "", core.ErrNotFound.WithDebug(err.Error())
		}
		return "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	return string(user.Status), nil
}
