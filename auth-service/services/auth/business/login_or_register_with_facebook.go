package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (b *business) LoginOrRegisterWithFacebook(ctx context.Context, userInfo *entity.OAuthFacebookUserInfo) (*entity.TokenResponse, error) {
	authData, err := b.repository.GetAuth(ctx, userInfo.Email)
	if err == nil && authData != nil {
		// Generate token using utility function
		return utils.GenerateToken(ctx, b.jwtProvider, authData.UserId)
	}

	splitName := func(fullName string) (string, string) {
		parts := strings.Fields(fullName)
		if len(parts) == 0 {
			return "", ""
		}

		firstName := parts[0]
		lastName := strings.Join(parts[1:], " ")

		return firstName, lastName
	}

	firstName, lastName := splitName(userInfo.Name)

	newUserId, err := b.userRepository.CreateUser(ctx,
		firstName,
		lastName,
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

			// Tạo auth record mới với userId đã tồn tại
			newAuth := entity.Auth{
				SQLModel:   core.SQLModel{},
				UserId:     userId,
				AuthType:   "facebook",
				Email:      userInfo.Email,
				FacebookId: userInfo.ID,
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
		SQLModel:   core.SQLModel{},
		UserId:     newUserId,
		AuthType:   "facebook",
		Email:      userInfo.Email,
		FacebookId: userInfo.ID,
	}

	if err := b.repository.AddNewAuth(ctx, &newAuth); err != nil {
		fmt.Printf("Error adding auth record: %v\n", err)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Generate token using utility function
	return utils.GenerateToken(ctx, b.jwtProvider, newUserId)
}
