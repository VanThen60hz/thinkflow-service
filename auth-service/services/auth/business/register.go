package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

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

	otp := core.GenerateOTP()

	// Use the utility function to send OTP email
	err = utils.SendOTPEmail(
		ctx,
		biz.redisClient,
		biz.emailService,
		data.Email,
		otp,
		common.EmailVerifyOTPSubject,
		"Email Verification",
		"Thanks for signing up! Please use the OTP below to verify your email:",
		"email verification",
	)
	if err != nil {
		return err
	}

	return nil
}
