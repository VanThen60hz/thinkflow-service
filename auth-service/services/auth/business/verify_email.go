package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) VerifyEmail(ctx context.Context, data *entity.EmailVerificationRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	auth, err := biz.ValidateEmailAndGetAuth(ctx, data.Email)
	if err != nil {
		return err
	}

	if err := biz.VerifyOTP(ctx, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	oldStatus, err := biz.userRepository.GetUserStatus(ctx, auth.UserId)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err := biz.userRepository.UpdateUserStatus(ctx, auth.UserId, "active"); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err := biz.DeleteOTP(ctx, data.Email, "verification:otp"); err != nil {
		if err := biz.userRepository.UpdateUserStatus(ctx, auth.UserId, oldStatus); err != nil {
			biz.CompensateUserCreation(ctx, auth.UserId)
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
