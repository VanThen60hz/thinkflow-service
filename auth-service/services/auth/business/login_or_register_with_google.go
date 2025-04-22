package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
)

func (biz *business) LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error) {
	return biz.ProcessOAuthLogin(
		ctx,
		userInfo.Email,
		userInfo.GivenName,
		userInfo.FamilyName,
		userInfo.ID,
		"google",
	)
}
