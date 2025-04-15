package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
)

func (b *business) ProcessFacebookCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error) {
	// The OAuth component will validate the state internally
	userInfo, err := b.oauthProvider.ProcessFacebookCallback(ctx, code, state)
	if err != nil {
		return nil, err
	}

	return b.LoginOrRegisterWithFacebook(ctx, &entity.OAuthFacebookUserInfo{
		ID:    userInfo.ID,
		Email: userInfo.Email,
		Name:  userInfo.FirstName + " " + userInfo.LastName,
	})
}
