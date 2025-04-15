package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *api) FacebookLoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		url := api.business.GetFacebookAuthURL()
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}
