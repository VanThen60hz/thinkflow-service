package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
)

func (b *business) ProcessGoogleCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error) {
	// The OAuth component will validate the state internally
	userInfo, err := b.oauthProvider.ProcessGoogleCallback(ctx, code, state)
	if err != nil {
		return nil, err
	}

	return b.LoginOrRegisterWithGoogle(ctx, &entity.OAuthGoogleUserInfo{
		ID:         userInfo.ID,
		Email:      userInfo.Email,
		GivenName:  userInfo.FirstName,
		FamilyName: userInfo.LastName,
	})
}
