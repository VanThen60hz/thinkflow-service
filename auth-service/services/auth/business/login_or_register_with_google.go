package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"
)

func (b *business) LoginOrRegisterWithGoogle(ctx context.Context, userInfo *entity.OAuthGoogleUserInfo) (*entity.TokenResponse, error) {
	return utils.ProcessOAuthLogin(
		ctx,
		b.repository,
		b.userRepository,
		b.jwtProvider,
		userInfo.Email,
		userInfo.GivenName,
		userInfo.FamilyName,
		userInfo.ID,
		"google",
	)
}
