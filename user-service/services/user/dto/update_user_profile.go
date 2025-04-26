package dto

import (
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type UpdateUserProfileRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Gender    *entity.Gender `json:"gender,omitempty"`
	AvatarId  *string        `json:"avatar_id,omitempty"`
}

func (req *UpdateUserProfileRequest) ToUserDataUpdate() (*entity.UserDataUpdate, error) {
	updateData := &entity.UserDataUpdate{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Gender:    req.Gender,
	}

	if req.AvatarId != nil {
		avatarUid, err := core.FromBase58(*req.AvatarId)
		if err != nil {
			return nil, err
		}
		avatarId := int(avatarUid.GetLocalID())
		updateData.Avatar = &avatarId
	}

	return updateData, nil
}
