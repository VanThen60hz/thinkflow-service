package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/gin-gonic/gin"
)

func (api *api) GoogleLoginHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		url := common.AppOAuth2Config.GoogleConfig.AuthCodeURL(oauthStateString)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
