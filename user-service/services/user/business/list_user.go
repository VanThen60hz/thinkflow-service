package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ListUsers(ctx context.Context, filter *entity.UserFilter, paging *core.Paging) ([]entity.User, error) {
	requesterVal := ctx.Value(common.RequesterKey)
	if requesterVal == nil {
		return nil, core.ErrUnauthorized.WithError("requester not found in context")
	}

	requester, ok := requesterVal.(core.Requester)
	if !ok {
		return nil, core.ErrInternalServerError.WithError("invalid requester type in context")
	}

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
			WithError("cannot access: user is not an admin or super admin.").
			WithDebug("requester role: " + string(requesterRole))
	}

	users, err := biz.userRepo.ListUsers(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot list users").
			WithDebug(err.Error())
	}

	avatarIds := make([]int, 0)
	for i := range users {
		if users[i].AvatarId > 0 {
			avatarIds = append(avatarIds, users[i].AvatarId)
		}
	}

	avatarMap := make(map[int]*core.Image)
	for _, avatarId := range avatarIds {
		img, err := biz.imageRepo.GetImageById(ctx, avatarId)
		if err != nil {
			return nil, core.ErrInternalServerError.
				WithError("cannot get user avatar").
				WithDebug(err.Error())
		}
		avatarMap[avatarId] = img
	}

	for i := range users {
		if users[i].AvatarId > 0 {
			if avatar, ok := avatarMap[users[i].AvatarId]; ok {
				users[i].Avatar = avatar
			}
		}
	}

	return users, nil
}
