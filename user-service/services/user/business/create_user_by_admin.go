package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type CreateUserData struct {
	FirstName  string            `json:"first_name,omitempty"`
	LastName   string            `json:"last_name,omitempty"`
	Email      string            `json:"email,omitempty"`
	Password   string            `json:"password,omitempty"`
	Phone      string            `json:"phone,omitempty"`
	AvatarId   *string           `json:"avatar_id,omitempty"`
	Gender     entity.Gender     `json:"gender,omitempty"`
	SystemRole entity.SystemRole `json:"role,omitempty"`
	Status     entity.Status     `json:"status,omitempty"`
}

func (biz *business) CreateUserByAdmin(ctx context.Context, data *CreateUserData) (*entity.User, error) {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	requesterDetail, err := biz.GetUserDetails(ctx, requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get requester status").
			WithDebug(err.Error())
	}

	requesterRole := requesterDetail.SystemRole
	if requesterRole != "admin" && requesterRole != "sadmin" {
		return nil, core.ErrForbidden.
			WithError("cannot access: user is not an admin or super admin").
			WithDebug("requester role: " + string(requesterRole))
	}

	userData := entity.NewUserForCreation(
		data.FirstName,
		data.LastName,
		data.Email,
	)

	if err := biz.userRepo.CreateNewUser(ctx, &userData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	updateData := &entity.UserDataUpdate{
		Phone:  &data.Phone,
		Gender: &data.Gender,
		Status: &data.Status,
	}

	if data.SystemRole != "" {
		updateData.SystemRole = &data.SystemRole
	}

	if data.AvatarId != nil {
		avatarUid, err := core.FromBase58(*data.AvatarId)
		if err != nil {
			return nil, err
		}
		avatarId := int(avatarUid.GetLocalID())
		updateData.Avatar = &avatarId
	}

	if err := updateData.Validate(); err != nil {
		return nil, core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.userRepo.UpdateUser(ctx, userData.Id, updateData); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	if err := biz.authRepo.RegisterWithUserId(ctx, int32(userData.Id), data.Email, data.Password); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	user, err := biz.userRepo.GetUserById(ctx, userData.Id)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	if user.AvatarId == 0 {
		return user, nil
	}

	image, err := biz.imageRepo.GetImageById(ctx, user.AvatarId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(entity.ErrCannotGetUser.Error()).
			WithDebug(err.Error())
	}

	user.Avatar = image

	return user, nil
}
