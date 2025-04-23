package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) DeleteNoteMemberHdl() func(*gin.Context) {
	return func(c *gin.Context) {
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

		if err := api.business.DeleteNoteMember(c.Request.Context(), int(noteId.GetLocalID()), int(userId.GetLocalID())); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
