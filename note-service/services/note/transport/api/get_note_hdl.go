package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (a *api) GetNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data, err := a.business.GetNoteById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()

		response := NewNoteResponse(data)

		c.JSON(http.StatusOK, core.ResponseData(response))
	}
}
