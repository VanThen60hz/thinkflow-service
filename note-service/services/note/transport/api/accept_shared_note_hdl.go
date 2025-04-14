package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) AcceptSharedNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		sharedToken := c.Param("token")

		note, err := api.business.AcceptSharedNote(ctx, sharedToken)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		note.Mask()

		c.JSON(http.StatusOK, core.ResponseData(gin.H{
			"message": "Accepted shared note successfully",
			"note":    note,
		}))
	}
}
