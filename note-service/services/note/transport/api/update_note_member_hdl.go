package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type UpdateNoteMemberRequest struct {
	Permission string `json:"permission" binding:"required"`
}

func (api *api) UpdateNoteMemberHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var req UpdateNoteMemberRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		noteIdStr := c.Param("note-id")
		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Invalid note ID"))
			return
		}

		userIdStr := c.Param("user-id")
		userId, err := core.FromBase58(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Invalid user ID"))
			return
		}

		if err := api.business.UpdateNoteMemberAccess(c.Request.Context(), int(noteId.GetLocalID()), int(userId.GetLocalID()), req.Permission); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
