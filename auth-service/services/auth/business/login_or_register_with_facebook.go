package business

import (
	"context"
	"strings"

	"thinkflow-service/services/auth/entity"
)

func (biz *business) LoginOrRegisterWithFacebook(ctx context.Context, userInfo *entity.OAuthFacebookUserInfo) (*entity.TokenResponse, error) {
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

	return biz.ProcessOAuthLogin(
		ctx,
		userInfo.Email,
		firstName,
		lastName,
		userInfo.ID,
		"facebook",
	)
}
