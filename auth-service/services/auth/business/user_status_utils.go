package business

import (
	"context"
	"fmt"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) CheckUserStatus(ctx context.Context, userId int, expectedStatus string) error {
	status, err := biz.userRepository.GetUserStatus(ctx, userId)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if status != expectedStatus {
		return core.ErrBadRequest.WithError(fmt.Sprintf("%s: %s", entity.ErrUserStatusNotMatch.Error(), expectedStatus))
	}

	return nil
}

func (biz *business) IsUserWaitingVerification(ctx context.Context, userId int) (bool, error) {
	status, err := biz.userRepository.GetUserStatus(ctx, userId)
	if err != nil {
		return false, core.ErrInternalServerError.WithDebug(err.Error())
	}

	return status == "waiting_verify", nil
}
