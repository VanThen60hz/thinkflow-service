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

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    response.AccessToken.Token,
			Path:     "/",
			Domain:   "",
			MaxAge:   604800,
			HttpOnly: true,
			Secure:   core.IsHTTPS(c),
			SameSite: http.SameSiteNoneMode,
		})

		c.JSON(http.StatusOK, core.ResponseData("Login successfully"))
	}
}
