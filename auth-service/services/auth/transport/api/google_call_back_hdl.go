package api

import (
	"net/http"
	"os"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GoogleCallbackHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		tokenResponse, err := api.business.ProcessGoogleCallback(c.Request.Context(), c.Query("code"), c.Query("state"))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		core.SetAccessTokenCookieWithDefaultPath(c, tokenResponse.AccessToken.Token)
		clientURL := os.Getenv("CLIENT_URL")
		if clientURL == "" {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError("CLIENT_URL is not set"))
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, clientURL+"/workspace/settings")
	}
}
