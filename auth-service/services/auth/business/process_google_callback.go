package business

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (b *business) ProcessGoogleCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error) {
	if err := b.ValidateOAuthState(state); err != nil {
		return nil, err
	}

	token, err := common.AppOAuth2Config.GoogleConfig.Exchange(ctx, code)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("code exchange failed").WithDebug(err.Error())
	}

	userInfoURL := os.Getenv("GOOGLE_USER_INFO_URL")
	resp, err := http.Get(userInfoURL + "?access_token=" + token.AccessToken)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed getting user info").WithDebug(err.Error())
	}
	defer resp.Body.Close()

	var userInfo entity.OAuthGoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed decoding user info").WithDebug(err.Error())
	}

	return b.LoginOrRegisterWithGoogle(ctx, &userInfo)
}
