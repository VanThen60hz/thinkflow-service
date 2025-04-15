package api

import (
	"net/http"
	"thinkflow-service/common"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (api *api) FacebookLoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		opts := []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam("auth_type", "rerequest"),
		}
		url := common.AppOAuth2Config.FacebookConfig.AuthCodeURL(oauthStateString, opts...)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}