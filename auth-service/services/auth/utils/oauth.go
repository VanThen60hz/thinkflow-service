package utils

import (
	"context"
	"errors"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func ProcessOAuthLogin(ctx context.Context, repository AuthRepository, userRepository UserRepository,
	jwtProvider common.JWTProvider, email, firstName, lastName, oauthID, oauthType string,
) (*entity.TokenResponse, error) {
	authData, err := repository.GetAuth(ctx, email)
	if err == nil && authData != nil {
		return GenerateToken(ctx, jwtProvider, authData.UserId)
	}

	newUserId, err := userRepository.CreateUser(ctx, firstName, lastName, email)
	if err != nil {
		if isDuplicateEmailError(err) {
			return handleExistingEmail(ctx, repository, userRepository, jwtProvider, email, oauthID, oauthType)
		}
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := createOAuthAuthRecord(newUserId, email, oauthID, oauthType)
	if err := repository.AddNewAuth(ctx, &newAuth); err != nil {
		CompensateUserCreation(ctx, userRepository, newUserId)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	token, err := GenerateToken(ctx, jwtProvider, newUserId)
	if err != nil {
		CompensateAuthCreation(ctx, repository, email)
		CompensateUserCreation(ctx, userRepository, newUserId)
		return nil, err
	}

	return token, nil
}

func isDuplicateEmailError(err error) bool {
	return errors.Is(err, entity.ErrDuplicateEmail)
}

func handleExistingEmail(ctx context.Context, repository AuthRepository, userRepository UserRepository,
	jwtProvider common.JWTProvider, email, oauthID, oauthType string,
) (*entity.TokenResponse, error) {
	userId, err := getUserIdByEmail(ctx, userRepository, email)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := createOAuthAuthRecord(userId, email, oauthID, oauthType)
	if err := repository.AddNewAuth(ctx, &newAuth); err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return GenerateToken(ctx, jwtProvider, userId)
}

func getUserIdByEmail(ctx context.Context, userRepository UserRepository, email string) (int, error) {
	return userRepository.GetUserIdByEmail(ctx, email)
}

func createOAuthAuthRecord(userId int, email, oauthID, oauthType string) entity.Auth {
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
