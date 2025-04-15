package api

import (
	"net/http"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ForgotPasswordHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.ForgotPasswordRequest

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		err := api.business.ForgotPassword(c.Request.Context(), &data)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}