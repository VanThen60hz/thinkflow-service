package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error)
	CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error
	UpdateUser(ctx context.Context, id int, data *entity.UserDataUpdate) error
}

type business struct {
	repository UserRepository
}

func NewBusiness(repository UserRepository) *business {
	return &business{repository: repository}
}

func (biz *business) GetUserProfile(ctx context.Context) (*entity.User, error) {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	user, err := biz.repository.GetUserById(ctx, requesterId)
	if err != nil {
		return nil, core.ErrUnauthorized.
			WithError(entity.ErrCannotGetUser.Error()).
			WithDebug(err.Error())
	}

	return user, nil
}

func (biz *business) GetUserDetails(ctx context.Context, id int) (*entity.User, error) {
	user, err := biz.repository.GetUserById(ctx, id)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrNotFound.
				WithError(entity.ErrCannotGetUser.Error()).
				WithDebug(err.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetUser.Error()).
			WithDebug(err.Error())
	}

	return user, nil
}

func (biz *business) GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error) {
	users, err := biz.repository.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, core.ErrNotFound.
			WithError(entity.ErrCannotGetUsers.Error()).
			WithDebug(err.Error())
	}

	return users, nil
}

func (biz *business) CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error {
	err := biz.repository.CreateNewUser(ctx, data)
	if err != nil {
		return core.ErrInternalServerError.
			WithError(entity.ErrCannotCreateUser.Error()).
			WithDebug(err.Error())
	}

	return nil
}

func (biz *business) UpdateUserProfile(ctx context.Context, data *entity.UserDataUpdate) error {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.repository.UpdateUser(ctx, requesterId, data); err != nil {
		return core.ErrInternalServerError.
			WithError("cannot update user profile").
			WithDebug(err.Error())
	}

	return nil
}
