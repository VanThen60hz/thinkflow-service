package business

import (
	"context"
	"fmt"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Get stored OTP from Redis
	key := fmt.Sprintf("otp:%s", data.Email)
	storedOTP, err := biz.redisClient.Get(ctx, key)
	if err != nil {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	// Verify OTP
	if storedOTP != data.OTP {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	// Generate new salt and hash password
	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err := biz.hasher.HashPassword(salt, data.NewPassword)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Update password in database
	if err := biz.repository.UpdatePassword(ctx, data.Email, salt, hashedPassword); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Delete OTP from Redis after successful password reset
	_ = biz.redisClient.Del(ctx, key)

	return nil
}
