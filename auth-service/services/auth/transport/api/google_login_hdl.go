package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *api) GoogleLoginHdl() func(c *gin.Context) {
	return func(c *gin.Context) {
		url := api.business.GetGoogleAuthURL()
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
