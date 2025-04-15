package api

import (
	"context"

	"thinkflow-service/services/auth/entity"

	sctx "github.com/VanThen60hz/service-context"
)

type Business interface {
	Login(ctx context.Context, data *entity.AuthEmailPassword) (*entity.TokenResponse, error)
	Register(ctx context.Context, data *entity.AuthRegister) error
	ForgotPassword(ctx context.Context, data *entity.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, data *entity.ResetPasswordRequest) error
	Logout(ctx context.Context, accessToken string) error
	VerifyEmail(ctx context.Context, data *entity.EmailVerificationRequest) error
	ResendVerificationOTP(ctx context.Context, data *entity.ResendOTPRequest) error
	// OAuth methods
	GetFacebookAuthURL() string
	GetGoogleAuthURL() string
	ProcessGoogleCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error)
	ProcessFacebookCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error)
	LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error)
	LoginOrRegisterWithFacebook(ctx context.Context, userInfo *entity.OAuthFacebookUserInfo) (*entity.TokenResponse, error)
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{serviceCtx: serviceCtx, business: business}
}
