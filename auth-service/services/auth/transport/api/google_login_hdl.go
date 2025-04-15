package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/gin-gonic/gin"
)

func (api *api) GoogleLoginHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		url := api.business.GetGoogleAuthURL(common.OAuthStateString)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
