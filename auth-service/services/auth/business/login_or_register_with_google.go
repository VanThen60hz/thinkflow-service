package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

func (b *business) LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error) {
	authData, err := b.repository.GetAuth(ctx, userInfo.Email)
	if err == nil && authData != nil {
		uid := core.NewUID(uint32(authData.UserId), 1, 1)
		sub := uid.String()
		tid := uuid.New().String()

		tokenStr, expSecs, err := b.jwtProvider.IssueToken(ctx, tid, sub)
		if err != nil {
			return nil, core.ErrInternalServerError.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
		}

		return &entity.TokenResponse{
			AccessToken: entity.Token{
				Token:     tokenStr,
				ExpiredIn: expSecs,
			},
		}, nil
	}

	newUserId, err := b.userRepository.CreateUser(ctx,
		userInfo.GivenName,
		userInfo.FamilyName,
		userInfo.Email,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "for key 'users.email'") {
			// Tạm thời sử dụng một cách khác để lấy user_id
			// TODO: Thay thế bằng GetUserIdByEmail khi có
			var userId int
			err = b.userRepository.GetDB().Table("users").
				Select("id").
				Where("email = ?", userInfo.Email).
				First(&userId).Error
			if err != nil {
				return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
			}

			newAuth := entity.Auth{
				SQLModel: core.SQLModel{},
				UserId:   userId,
				AuthType: "google",
				Email:    userInfo.Email,
				GoogleId: userInfo.ID,
			}

			if err := b.repository.AddNewAuth(ctx, &newAuth); err != nil {
				fmt.Printf("Error adding auth record: %v\n", err)
				return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
			}

			uid := core.NewUID(uint32(userId), 1, 1)
			sub := uid.String()
			tid := uuid.New().String()

			tokenStr, expSecs, err := b.jwtProvider.IssueToken(ctx, tid, sub)
			if err != nil {
				return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
			}

			return &entity.TokenResponse{
				AccessToken: entity.Token{
					Token:     tokenStr,
					ExpiredIn: expSecs,
				},
			}, nil
		}
		fmt.Printf("Error creating user: %v\n", err)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := entity.Auth{
		SQLModel: core.SQLModel{},
		UserId:   newUserId,
		AuthType: "google",
		Email:    userInfo.Email,
		GoogleId: userInfo.ID,
	}

	if err := b.repository.AddNewAuth(ctx, &newAuth); err != nil {
		fmt.Printf("Error adding auth record: %v\n", err)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	uid := core.NewUID(uint32(newUserId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := b.jwtProvider.IssueToken(ctx, tid, sub)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return &entity.TokenResponse{
		AccessToken: entity.Token{
			Token:     tokenStr,
			ExpiredIn: expSecs,
		},
	}, nil
}
