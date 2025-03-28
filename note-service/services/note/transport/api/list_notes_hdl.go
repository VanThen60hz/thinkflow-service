package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := core.GetRequester(c).GetSubject()

		rp.Paging.Process()
		rp.UserId = &requester

		note, err := api.business.ListNotes(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		for i := range note {
			note[i].Mask()
		}

		c.JSON(http.StatusOK, core.SuccessResponse(note, rp.Paging, rp.Filter))
	}
}
