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

func (biz *business) ResendVerificationOTP(ctx context.Context, data *entity.ResendOTPRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Check if email exists
	auth, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Check user status using RPC method
	status, err := biz.userRepository.GetUserStatus(ctx, auth.UserId)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if status != "waiting_verify" {
		return core.ErrBadRequest.WithError("User is already verified")
	}

	// Generate new OTP
	otp := core.GenerateOTP()

	// Update OTP in Redis with 10 minutes expiration
	key := fmt.Sprintf("verification:otp:%s", data.Email)
	if err := biz.redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	emailData := emailc.OTPMailData{
		Title:         "Email Verification",
		UserEmail:     data.Email,
		MessageIntro:  "Thanks for signing up! Please use the OTP below to verify your email:",
		OTP:           otp,
		OTPTypeDesc:   "email verification",
		ExpireMinutes: 10,
	}

	// Send verification OTP via email
	if err := biz.emailService.SendGenericOTP(ctx, data.Email, common.EmailVerifyOTPSubject, emailData); err != nil {
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
