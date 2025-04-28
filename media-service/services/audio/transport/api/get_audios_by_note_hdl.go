package api

import (
	"net/http"
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetAudiosByNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteId, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		data, err := api.business.GetAudiosByNoteId(ctx, int(noteId.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for _, audio := range data {
			audio.Mask()
		}

		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}
