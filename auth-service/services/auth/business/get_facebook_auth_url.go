package business

import (
	"thinkflow-service/common"

	"golang.org/x/oauth2"
)

func (b *business) GetFacebookAuthURL(state string) string {
	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("auth_type", "rerequest"),
	}
	return common.AppOAuth2Config.FacebookConfig.AuthCodeURL(state, opts...)
}
