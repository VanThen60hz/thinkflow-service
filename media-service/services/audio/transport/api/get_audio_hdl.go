package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		audioId, err := core.FromBase58(c.Param("audio-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		data, err := api.business.GetAudioById(ctx, int(audioId.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
