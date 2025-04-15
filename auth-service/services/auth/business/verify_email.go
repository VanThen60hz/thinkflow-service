package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) VerifyEmail(ctx context.Context, data *entity.EmailVerificationRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	auth, err := utils.ValidateEmailAndGetAuth(ctx, biz.repository, data.Email)
	if err != nil {
		return err
	}

	if err := utils.VerifyOTP(ctx, biz.redisClient, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	if err := biz.userRepository.UpdateUserStatus(ctx, auth.UserId, "active"); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	utils.DeleteOTP(ctx, biz.redisClient, data.Email, "verification:otp")

	return nil
}
