package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetImageHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("image-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data, err := api.business.GetImageById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
