package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.NoteDataCreation

		if err := c.ShouldBind(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if data.Title == "" {
			data.Title = "New Note"
		}

		// Set requester to context
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		// we can set user_id directly to data creation, but I don't recommend it
		// uid, _ := core.FromBase58(requester.GetSubject())
		// data.UserId = int(uid.GetLocalID())

		if err := api.business.CreateNewNote(ctx, &data); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
