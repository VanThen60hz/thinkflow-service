package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) FacebookCallbackHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		tokenResponse, err := api.business.ProcessFacebookCallback(c.Request.Context(), c.Query("code"), c.Query("state"))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		core.SetAccessTokenCookieWithDefaultPath(c, tokenResponse.AccessToken.Token)
		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}
