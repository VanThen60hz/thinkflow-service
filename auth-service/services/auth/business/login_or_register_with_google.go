package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (b *business) LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error) {
	authData, err := b.repository.GetAuth(ctx, userInfo.Email)
	if err == nil && authData != nil {
		// Generate token using utility function
		return utils.GenerateToken(ctx, b.jwtProvider, authData.UserId)
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

			// Generate token using utility function
			return utils.GenerateToken(ctx, b.jwtProvider, userId)
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

	// Generate token using utility function
	return utils.GenerateToken(ctx, b.jwtProvider, newUserId)
}
