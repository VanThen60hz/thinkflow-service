package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.userRepo.UpdateUser(ctx, requesterId, data); err != nil {
		return core.ErrInternalServerError.
			WithError("cannot update user profile").
			WithDebug(err.Error())
	}

	return nil
}
