package business

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	AddNewAuth(ctx context.Context, data *entity.Auth) error
	GetAuth(ctx context.Context, email string) (*entity.Auth, error)
	UpdatePassword(ctx context.Context, email, salt, hashedPassword string) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, firstName, lastName, email string) (newId int, err error)
	GetUserIdByEmail(ctx context.Context, email string) (int, error)
	GetDB() *gorm.DB
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type business struct {
	repository     AuthRepository
	userRepository UserRepository
	jwtProvider    common.JWTProvider
	hasher         Hasher
	redisClient    *common.RedisClient
	emailService   *common.EmailService
}

func NewBusiness(repository AuthRepository, userRepository UserRepository,
	jwtProvider common.JWTProvider, hasher Hasher,
	redisClient *common.RedisClient, emailService *common.EmailService,
) *business {
	return &business{
		repository:     repository,
		userRepository: userRepository,
		jwtProvider:    jwtProvider,
		hasher:         hasher,
		redisClient:    redisClient,
		emailService:   emailService,
	}
}

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

	uid := core.NewUID(uint32(authData.UserId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := biz.jwtProvider.IssueToken(ctx, tid, sub)
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

func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	_, err := biz.repository.GetAuth(ctx, data.Email)

	if err == nil {
		return core.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != core.ErrRecordNotFound {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	passHashed, err := biz.hasher.HashPassword(salt, data.Password)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, passHashed)

	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	return nil
}

func (biz *business) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	// Check if token is blacklisted
	blacklistKey := fmt.Sprintf("blacklist:token:%s", accessToken)
	_, err := biz.redisClient.Get(ctx, blacklistKey)
	if err == nil {
		// Token is blacklisted
		return nil, core.ErrUnauthorized.WithError("token has been revoked")
	}

	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)
	if err != nil {
		return nil, core.ErrUnauthorized.WithDebug(err.Error())
	}

	return claims, nil
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	otp := ""
	for i := 0; i < 6; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}
	return otp
}

func (biz *business) ForgotPassword(ctx context.Context, data *entity.ForgotPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Check if email exists
	_, err := biz.repository.GetAuth(ctx, data.Email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Generate OTP
	otp := generateOTP()

	// Store OTP in Redis with 10 minutes expiration
	key := fmt.Sprintf("otp:%s", data.Email)
	if err := biz.redisClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Send OTP via email
	if err := biz.emailService.SendOTP(data.Email, otp); err != nil {
		// Delete OTP from Redis if email fails
		_ = biz.redisClient.Del(ctx, key)
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}

func (biz *business) ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Get stored OTP from Redis
	key := fmt.Sprintf("otp:%s", data.Email)
	storedOTP, err := biz.redisClient.Get(ctx, key)
	if err != nil {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	// Verify OTP
	if storedOTP != data.OTP {
		return core.ErrBadRequest.WithError(entity.ErrInvalidOrExpiredOTP.Error())
	}

	// Generate new salt and hash password
	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err := biz.hasher.HashPassword(salt, data.NewPassword)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Update password in database
	if err := biz.repository.UpdatePassword(ctx, data.Email, salt, hashedPassword); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	// Delete OTP from Redis after successful password reset
	_ = biz.redisClient.Del(ctx, key)

	return nil
}

func (b *business) LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error) {
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
	fmt.Printf("Creating new user with info: email=%s, given_name=%s, family_name=%s\n",
		userInfo.Email, userInfo.GivenName, userInfo.FamilyName)

	newUserId, err := b.userRepository.CreateUser(ctx,
		userInfo.GivenName,
		userInfo.FamilyName,
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

	// Tạo auth record mới
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

	// Tạo auth record mới
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

func (biz *business) Logout(ctx context.Context, accessToken string) error {
	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)
	if err != nil {
		return core.ErrUnauthorized.WithDebug(err.Error())
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	ttl := time.Until(expTime.Time)
	if ttl <= 0 {
		return nil
	}

	blacklistKey := fmt.Sprintf("blacklist:token:%s", accessToken)
	err = biz.redisClient.Set(ctx, blacklistKey, "revoked", ttl)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
