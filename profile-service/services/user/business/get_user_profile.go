package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) GetUserProfile(ctx context.Context) (*entity.User, error) {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	user, err := biz.userRepo.GetUserById(ctx, requesterId)
	if err != nil {
		return nil, core.ErrUnauthorized.
			WithError(entity.ErrCannotGetUser.Error()).
			WithDebug(err.Error())
	}

	// Get extra infos: User
	image, err := biz.imageRepo.GetImageById(ctx, user.AvatarId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetUser.Error()).
			WithDebug(err.Error())
	}

	user.Avatar = image

	return user, nil
}
