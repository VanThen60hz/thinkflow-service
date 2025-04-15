package utils

import (
	"context"
	"fmt"
	"time"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/core"
)

// SendOTPEmail sends an OTP email to the user
func SendOTPEmail(ctx context.Context, redisClient redisc.Redis, emailService emailc.Email, email, otp, subject, title, messageIntro, otpTypeDesc string) error {
	key := fmt.Sprintf("verification:otp:%s", email)
	if err := redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	emailData := emailc.OTPMailData{
		Title:         title,
		UserEmail:     email,
		MessageIntro:  messageIntro,
		OTP:           otp,
		OTPTypeDesc:   otpTypeDesc,
		ExpireMinutes: 10,
	}

	if err := emailService.SendGenericOTP(ctx, email, subject, emailData); err != nil {
		_ = redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}

// VerifyOTP checks if the provided OTP matches the stored OTP
func VerifyOTP(ctx context.Context, redisClient redisc.Redis, email, otp, keyPrefix string) error {
	key := fmt.Sprintf("%s:%s", keyPrefix, email)
	storedOTP, err := redisClient.Get(ctx, key)

	if err != nil {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	if storedOTP != otp {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	return nil
}

// DeleteOTP removes the OTP from Redis
func DeleteOTP(ctx context.Context, redisClient redisc.Redis, email, keyPrefix string) {
	key := fmt.Sprintf("%s:%s", keyPrefix, email)
	_ = redisClient.Del(ctx, key)
}
