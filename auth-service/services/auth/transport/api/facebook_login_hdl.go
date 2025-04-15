package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/gin-gonic/gin"
)

func (api *api) FacebookLoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		url := api.business.GetFacebookAuthURL(common.OAuthStateString)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
