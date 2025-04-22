package business

import (
	"context"
	"fmt"
	"time"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) SendOTPEmail(ctx context.Context, email, otp, subject, title, messageIntro, otpTypeDesc string) error {
	key := fmt.Sprintf("verification:otp:%s", email)
	if err := biz.redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
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

	if err := biz.emailService.SendGenericOTP(ctx, email, subject, emailData); err != nil {
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}

func (biz *business) VerifyOTP(ctx context.Context, email, otp, keyPrefix string) error {
	key := fmt.Sprintf("%s:%s", keyPrefix, email)
	storedOTP, err := biz.redisClient.Get(ctx, key)
	if err != nil {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	if storedOTP != otp {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	return nil
}

func (biz *business) DeleteOTP(ctx context.Context, email, keyPrefix string) error {
	key := fmt.Sprintf("%s:%s", keyPrefix, email)
	err := biz.redisClient.Del(ctx, key)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}
	return nil
}
