package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) LogoutHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("accessToken")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithError("missing access token in cookie"))
			return
		}

		if token == "" {
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithError("empty access token"))
			return
		}

		err = api.business.Logout(c.Request.Context(), token)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		core.ClearAccessTokenCookieWithDefaultPath(c)
		c.JSON(http.StatusOK, core.ResponseData("Logout successfully"))
	}
}
