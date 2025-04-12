package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListNoteMembersHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var paging core.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		noteIdStr := c.Param("note-id")
		noteId, err := core.FromBase58(noteIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Invalid note ID"))
			return
		}

		users, err := api.business.ListNoteMembersById(c.Request.Context(), int(noteId.GetLocalID()), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(users, paging, nil))
	}
}
