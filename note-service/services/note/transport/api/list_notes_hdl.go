package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListNotesHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		requesterSubject := requester.GetSubject()

		rp.Paging.Process()
		rp.UserId = &requesterSubject

		notes, err := api.business.ListNotes(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range notes {
			notes[i].Mask()
		}

		response := ListNoteResponse{
			Data:   make([]NoteResponse, len(notes)),
			Paging: rp.Paging,
		}

		for i := range notes {
			response.Data[i] = *NewNoteResponse(&notes[i])
		}

		c.JSON(http.StatusOK, core.SuccessResponse(response.Data, response.Paging, rp.Filter))
	}
}
