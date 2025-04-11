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

func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	_, err := biz.repository.GetAuth(ctx, data.Email)

	if err == nil {
		return core.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != core.ErrRecordNotFound {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Create user with waiting_verify status
	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	passHashed, err := biz.hasher.HashPassword(salt, data.Password)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, passHashed)

	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Generate OTP
	otp := core.GenerateOTP()

	// Store OTP in Redis with 10 minutes expiration
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
		// Delete OTP from Redis if email fails
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
