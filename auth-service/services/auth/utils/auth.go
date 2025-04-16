package utils

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/component/emailc"
	"github.com/VanThen60hz/service-context/component/redisc"
	"github.com/VanThen60hz/service-context/core"
)

func ValidateEmailAndGetAuth(ctx context.Context, repository AuthRepository, email string) (*entity.Auth, error) {
	auth, err := repository.GetAuth(ctx, email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return nil, core.ErrInternalServerError.WithDebug(err.Error())
	}
	return auth, nil
}

func ProcessPassword(hasher Hasher, password string) (salt, hashedPassword string, err error) {
	salt, err = hasher.RandomStr(16)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err = hasher.HashPassword(salt, password)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	return salt, hashedPassword, nil
}

func ValidateRegistrationInput(ctx context.Context, repository AuthRepository, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	_, err := repository.GetAuth(ctx, data.Email)
	if err == nil {
		return core.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != core.ErrRecordNotFound {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return nil
}

func SendVerificationEmail(ctx context.Context, redisClient redisc.Redis, emailService emailc.Email, email string) error {
	otp := core.GenerateOTP()
	return SendOTPEmail(
		ctx,
		redisClient,
		emailService,
		email,
		otp,
		common.EmailVerifyOTPSubject,
		"Email Verification",
		"Thanks for signing up! Please use the OTP below to verify your email:",
		"email verification",
	)
}
