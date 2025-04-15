package business

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ForgotPassword(ctx context.Context, data *entity.ForgotPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	_, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	otp := core.GenerateOTP()

	err = utils.SendOTPEmail(
		ctx,
		biz.redisClient,
		biz.emailService,
		data.Email,
		otp,
		common.EmailResetPasswordSubject,
		"Password Reset Request",
		"We received a request to reset your password for your ThinkFlow account. To proceed with the password reset, please use the following One-Time Password (OTP)",
		"email reset password",
	)
	if err != nil {
		return err
	}

	return nil
}
