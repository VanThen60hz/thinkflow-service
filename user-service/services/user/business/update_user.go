package business

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) UpdateUser(ctx context.Context, userId int, data *entity.UserDataUpdate) error {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	requesterDetail, err := biz.GetUserDetails(ctx, requesterId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError("cannot get requester status").
			WithDebug(err.Error())
	}

	requesterRole := requesterDetail.SystemRole
	if requesterRole != "admin" && requesterRole != "sadmin" {
		return core.ErrForbidden.
			WithError("cannot access: user is not an admin or super admin").
			WithDebug("requester role: " + string(requesterRole))
	}

	targetUser, err := biz.GetUserDetails(ctx, userId)
	if err != nil {
		return core.ErrInternalServerError.
			WithError("cannot get user").
			WithDebug(err.Error())
	}

	if targetUser.SystemRole == "admin" || targetUser.SystemRole == "sadmin" {
		return core.ErrForbidden.
			WithError("cannot update admin or super admin users")
	}

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.userRepo.UpdateUser(ctx, userId, data); err != nil {
		return core.ErrInternalServerError.
			WithError("cannot update user").
			WithDebug(err.Error())
	}

	return nil
}
