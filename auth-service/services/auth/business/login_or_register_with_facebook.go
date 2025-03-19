package business

import (
	"context"
	"fmt"
	"strings"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

func (b *business) LoginOrRegisterWithFacebook(ctx context.Context, userInfo *entity.OAuthFacebookUserInfo) (*entity.TokenResponse, error) {
	// Kiểm tra xem user đã tồn tại bằng email trong bảng auths
	authData, err := b.repository.GetAuth(ctx, userInfo.Email)
	if err == nil && authData != nil {
		// User đã tồn tại trong cả users và auths -> login
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

	// Tạo user mới trong user service
	fmt.Printf("Creating new user with info: email=%s, name=%s\n",
		userInfo.Email, userInfo.Name)

	splitName := func(fullName string) (string, string) {
		parts := strings.Fields(fullName)
		if len(parts) == 0 {
			return "", ""
		}

		firstname := parts[0]
		lastname := strings.Join(parts[1:], " ")

		return firstname, lastname
	}

	firstname, lastname := splitName(userInfo.Name)

	newUserId, err := b.userRepository.CreateUser(ctx,
		firstname,
		lastname,
		userInfo.Email,
	)
	if err != nil {
		// Nếu email đã tồn tại trong users, chúng ta sẽ tạo auth record cho user đó
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

			// Tạo token cho user
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

	fmt.Printf("Created new user with ID: %d\n", newUserId)

	newAuth := entity.Auth{
		SQLModel: core.SQLModel{},
		UserId:   newUserId,
		AuthType: "google",
		Email:    userInfo.Email,
		GoogleId: userInfo.ID,
	}

	fmt.Printf("Creating auth record: %+v\n", newAuth)

	if err := b.repository.AddNewAuth(ctx, &newAuth); err != nil {
		fmt.Printf("Error adding auth record: %v\n", err)
		return nil, core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Tạo token cho user mới
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
