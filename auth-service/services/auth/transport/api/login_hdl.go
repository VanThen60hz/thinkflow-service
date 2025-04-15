package api

import (
	"net/http"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) LoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		response, err := api.business.Login(c.Request.Context(), &data)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		core.SetAccessTokenCookieWithDefaultPath(c, response.AccessToken.Token)
		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}
