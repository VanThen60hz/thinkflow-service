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

	salt, hashedPassword, err := utils.ProcessPassword(biz.hasher, data.Password)
	if err != nil {
		return err
	}

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	defer func() {
		if err != nil {
			utils.CompensateUserCreation(ctx, biz.userRepository, newUserId)
		}
	}()

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, hashedPassword)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		err = core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
		return err
	}

	defer func() {
		if err != nil {
			utils.CompensateAuthCreation(ctx, biz.repository, data.Email)
		}
	}()

	otp := core.GenerateOTP()
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
