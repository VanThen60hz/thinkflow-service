package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := utils.VerifyOTP(ctx, biz.redisClient, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	auth, err := utils.ValidateEmailAndGetAuth(ctx, biz.repository, data.Email)
	if err != nil {
		return err
	}

	oldSalt := auth.Salt
	oldPassword := auth.Password

	salt, hashedPassword, err := utils.ProcessPassword(biz.hasher, data.NewPassword)
	if err != nil {
		return err
	}

	if err := biz.repository.UpdatePassword(ctx, data.Email, salt, hashedPassword); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err := utils.DeleteOTP(ctx, biz.redisClient, data.Email, "otp"); err != nil {
		if err := biz.repository.UpdatePassword(ctx, data.Email, oldSalt, oldPassword); err != nil {
			utils.CompensateAuthCreation(ctx, biz.repository, data.Email)
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
