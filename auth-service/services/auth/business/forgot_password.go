package business

import (
	"context"
	"fmt"
	"time"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ForgotPassword(ctx context.Context, data *entity.ForgotPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Check if email exists
	_, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Generate OTP
	otp := generateOTP()

	// Store OTP in Redis with 10 minutes expiration
	key := fmt.Sprintf("otp:%s", data.Email)
	if err := biz.redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Send OTP via email
	if err := biz.emailService.SendOTP(data.Email, otp); err != nil {
		// Delete OTP from Redis if email fails
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
