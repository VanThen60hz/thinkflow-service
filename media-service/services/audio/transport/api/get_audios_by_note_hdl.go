package api

import (
	"net/http"

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

		data, err := api.business.GetAudiosByNoteId(c.Request.Context(), int(noteId.GetLocalID()))
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
