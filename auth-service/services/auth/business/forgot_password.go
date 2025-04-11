package business

import (
	"context"
	"fmt"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
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
	otp := core.GenerateOTP()

	// Store OTP in Redis with 10 minutes expiration
	key := fmt.Sprintf("otp:%s", data.Email)
	if err := biz.redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	emailData := emailc.OTPMailData{
		Title:         "Password Reset Request",
		UserEmail:     data.Email,
		MessageIntro:  "We received a request to reset your password for your ThinkFlow account. To proceed with the password reset, please use the following One-Time Password (OTP)",
		OTP:           otp,
		OTPTypeDesc:   "email reset password",
		ExpireMinutes: 10,
	}

	// Send OTP via email
	if err := biz.emailService.SendGenericOTP(ctx, data.Email, common.EmailResetPasswordSubject, emailData); err != nil {
		// Delete OTP from Redis if email fails
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
