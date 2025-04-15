package business

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (b *business) ProcessFacebookCallback(ctx context.Context, code string, state string) (*entity.TokenResponse, error) {
	if err := b.ValidateOAuthState(state); err != nil {
		return nil, err
	}

	token, err := common.AppOAuth2Config.FacebookConfig.Exchange(ctx, code)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("code exchange failed").WithDebug(err.Error())
	}

	resp, err := http.Get("https://graph.facebook.com/v20.0/me?fields=id,name,email&access_token=" + token.AccessToken)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed getting user info").WithDebug(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("failed reading response").WithDebug(err.Error())
	}

	var userInfo entity.OAuthFacebookUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, core.ErrInternalServerError.WithError("failed parsing user info").WithDebug(err.Error())
	}

	return b.LoginOrRegisterWithFacebook(ctx, &userInfo)
}
