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
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("note-id parameter is required"))
			return
		}

		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError("invalid note-id format"))
			return
		}

		text, err := api.business.GetTextByNoteId(c.Request.Context(), int(noteId.GetLocalID()))
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		text.Mask()

		c.JSON(http.StatusOK, core.SuccessResponse(text, nil, nil))
	}
}
