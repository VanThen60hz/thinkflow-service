package api

import (
	"net/http"
	"thinkflow-service/common"
	
	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetTextByNoteIdHdl() func(*gin.Context) {
	return func(c *gin.Context) {
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

		// Set requester to context
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		text, err := api.business.GetTextByNoteId(ctx, int(noteId.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		text.Mask()

		c.JSON(http.StatusOK, core.SuccessResponse(text, nil, nil))
	}
}
