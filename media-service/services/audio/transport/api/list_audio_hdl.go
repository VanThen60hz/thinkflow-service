package api

import (
	"net/http"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListAudiosHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		noteIdStr := c.Param("note-id")
		if noteIdStr == "" {
			noteIdStr = c.Query("note-id")
		}

		if noteIdStr == "" {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("note-id parameter is required"))
			return
		}

		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid note-id format"))
			return
		}

		if rp.Filter.NoteID == nil {
			rp.Filter.NoteID = new(int64)
		}
		*rp.Filter.NoteID = int64(noteId.GetLocalID())

		rp.Paging.Process()

		audios, err := api.business.ListAudios(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range audios {
			audios[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(audios, rp.Paging, rp.Filter))
	}
}
