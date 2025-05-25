package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) AcceptSharedNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		sharedToken := c.Param("token")

		note, alreadyAccessed, err := api.business.AcceptSharedNote(ctx, sharedToken)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		note.Mask()

		message := "Accepted shared note successfully"
		if alreadyAccessed {
			message = "You already have access to this note"
		}

		c.JSON(http.StatusOK, core.ResponseData(gin.H{
			"message":          message,
			"note":             note,
			"already_accessed": alreadyAccessed,
		}))
	}
}
