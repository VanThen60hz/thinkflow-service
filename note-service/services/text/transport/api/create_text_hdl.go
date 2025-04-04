package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) CreateTextHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		noteId, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var data entity.TextDataCreation

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data.NoteID = int64(noteId.GetLocalID())

		// Set requester to context
		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		// we can set user_id directly to data creation, but I don't recommend it
		// uid, _ := core.FromBase58(requester.GetSubject())
		// data.UserId = int(uid.GetLocalID())

		if err := api.business.CreateNewText(ctx, &data); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		data.Mask()
		c.JSON(http.StatusOK, core.ResponseData(data.FakeId))
	}
}
