package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.VerifyOTP(ctx, data.Email, data.OTP, "verification:otp"); err != nil {
		return err
	}

	auth, err := biz.ValidateEmailAndGetAuth(ctx, data.Email)
	if err != nil {
		return err
	}

	oldSalt := auth.Salt
	oldPassword := auth.Password

	salt, hashedPassword, err := biz.ProcessPassword(data.NewPassword)
	if err != nil {
		return err
	}

	if err := biz.repository.UpdatePassword(ctx, data.Email, salt, hashedPassword); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	if err := biz.DeleteOTP(ctx, data.Email, "otp"); err != nil {
		if err := biz.repository.UpdatePassword(ctx, data.Email, oldSalt, oldPassword); err != nil {
			biz.CompensateAuthCreation(ctx, data.Email)
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
