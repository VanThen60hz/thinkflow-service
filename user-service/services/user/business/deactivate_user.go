package business

import (
	"context"
	"time"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
)

type DeactivateUserResult struct {
	Id            string    `json:"id"`
	Status        string    `json:"status"`
	DeactivatedAt time.Time `json:"deactivated_at"`
}

func (biz *business) DeactivateUser(ctx context.Context, userId int) (*DeactivateUserResult, error) {
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

	targetUser, err := biz.GetUserDetails(ctx, userId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get user").
			WithDebug(err.Error())
	}

	if targetUser.SystemRole == "admin" || targetUser.SystemRole == "sadmin" {
		return nil, core.ErrForbidden.
			WithError("cannot deactivate admin or super admin users")
	}

	status := entity.StatusBanned
	updateData := &entity.UserDataUpdate{
		Status: &status,
	}

	if err := biz.userRepo.UpdateUser(ctx, userId, updateData); err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot deactivate user").
			WithDebug(err.Error())
	}

	updatedUser, err := biz.GetUserDetails(ctx, userId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get updated user data").
			WithDebug(err.Error())
	}

	updatedUser.Mask()

	return &DeactivateUserResult{
		Id:            updatedUser.FakeId.String(),
		Status:        string(entity.StatusBanned),
		DeactivatedAt: time.Now(),
	}, nil
}
