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

	// Check if email exists in auths
	auth, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Check OTP using utility function
	if err := utils.VerifyOTP(ctx, biz.redisClient, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	// Update user status to active using the RPC method
	if err := biz.userRepository.UpdateUserStatus(ctx, auth.UserId, "active"); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Delete OTP using utility function
	utils.DeleteOTP(ctx, biz.redisClient, data.Email, "verification:otp")

	return nil
}
