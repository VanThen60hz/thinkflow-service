package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, core.ErrBadRequest.WithError(err.Error())
	}

	authData, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		return nil, core.ErrBadRequest.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
	}

	if !biz.hasher.CompareHashPassword(authData.Password, authData.Salt, data.Password) {
		return nil, core.ErrBadRequest.WithError(entity.ErrLoginFailed.Error())
	}

	// Check user status using utility function
	isWaiting, err := utils.IsUserWaitingVerification(ctx, biz.userRepository, authData.UserId)
	if err != nil {
		return nil, err
	}

	if isWaiting {
		return nil, core.ErrForbidden.WithError("Email address has not been verified. Please check your email for the verification code.")
	}

	// Generate token using utility function
	return utils.GenerateToken(ctx, biz.jwtProvider, authData.UserId)
}
