package api

import (
	"net/http"
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetAttachmentsByNoteIDHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteID, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		attachments, err := api.business.GetAttachmentsByNoteID(ctx, int64(noteID.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		for i := range attachments {
			attachments[i].Mask()
		}

		c.JSON(http.StatusOK, core.ResponseData(attachments))
	}
}