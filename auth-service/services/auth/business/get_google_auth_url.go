package business

import "thinkflow-service/common"

func (b *business) GetGoogleAuthURL(state string) string {
	return common.AppOAuth2Config.GoogleConfig.AuthCodeURL(state)
}
