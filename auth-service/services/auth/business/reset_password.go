package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Verify OTP using utility function
	if err := utils.VerifyOTP(ctx, biz.redisClient, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err := biz.hasher.HashPassword(salt, data.NewPassword)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err := biz.repository.UpdatePassword(ctx, data.Email, salt, hashedPassword); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Delete OTP using utility function
	utils.DeleteOTP(ctx, biz.redisClient, data.Email, "otp")

	return nil
}
