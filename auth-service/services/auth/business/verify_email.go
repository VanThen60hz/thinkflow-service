package business

import (
	"context"
	"fmt"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) VerifyEmail(ctx context.Context, data *entity.EmailVerificationRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Check if email exists in auths
	auth, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Check OTP from Redis
	key := fmt.Sprintf("verification:otp:%s", data.Email)
	storedOTP, err := biz.redisClient.Get(ctx, key)
	if err != nil {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	if storedOTP != data.OTP {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	// Update user status to active using the RPC method
	if err := biz.userRepository.UpdateUserStatus(ctx, auth.UserId, "active"); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Delete OTP from Redis
	_ = biz.redisClient.Del(ctx, key)

	return nil
}
