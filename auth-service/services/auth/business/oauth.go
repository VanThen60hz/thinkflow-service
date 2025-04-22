package business

import (
	"context"
	"errors"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ProcessOAuthLogin(ctx context.Context, email, firstName, lastName, oauthID, oauthType string) (*entity.TokenResponse, error) {
	authData, err := biz.repository.GetAuth(ctx, email)
	if err == nil && authData != nil {
		return biz.GenerateToken(ctx, authData.UserId)
	}

	newUserId, err := biz.userRepository.CreateUser(ctx, firstName, lastName, email)
	if err != nil {
		if biz.isDuplicateEmailError(err) {
			return biz.handleExistingEmail(ctx, email, oauthID, oauthType)
		}
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := biz.createOAuthAuthRecord(newUserId, email, oauthID, oauthType)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		biz.CompensateUserCreation(ctx, newUserId)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	token, err := biz.GenerateToken(ctx, newUserId)
	if err != nil {
		biz.CompensateAuthCreation(ctx, email)
		biz.CompensateUserCreation(ctx, newUserId)
		return nil, err
	}

	return token, nil
}

func (biz *business) isDuplicateEmailError(err error) bool {
	return errors.Is(err, entity.ErrDuplicateEmail)
}

func (biz *business) handleExistingEmail(ctx context.Context, email, oauthID, oauthType string) (*entity.TokenResponse, error) {
	userId, err := biz.getUserIdByEmail(ctx, email)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := biz.createOAuthAuthRecord(userId, email, oauthID, oauthType)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return biz.GenerateToken(ctx, userId)
}

func (biz *business) getUserIdByEmail(ctx context.Context, email string) (int, error) {
	return biz.userRepository.GetUserIdByEmail(ctx, email)
}

func (biz *business) createOAuthAuthRecord(userId int, email, oauthID, oauthType string) entity.Auth {
	newAuth := entity.Auth{
		SQLModel: core.SQLModel{},
		UserId:   userId,
		AuthType: oauthType,
		Email:    email,
	}

	if oauthType == "facebook" {
		newAuth.FacebookId = oauthID
	} else if oauthType == "google" {
		newAuth.GoogleId = oauthID
	}

	return newAuth
}
