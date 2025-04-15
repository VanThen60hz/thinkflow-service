package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ResendVerificationOTP(ctx context.Context, data *entity.ResendOTPRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	auth, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	isWaiting, err := utils.IsUserWaitingVerification(ctx, biz.userRepository, auth.UserId)
	if err != nil {
		return err
	}

	if !isWaiting {
		return core.ErrBadRequest.WithError(entity.ErrUserAlreadyVerified.Error())
	}

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
