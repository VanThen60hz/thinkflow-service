package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("audio-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data, err := api.business.GetAudioById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
