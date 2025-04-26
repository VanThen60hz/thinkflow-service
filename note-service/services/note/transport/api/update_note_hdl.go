package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type NoteDataUpdateRequest struct {
	Title     *string `json:"title" gorm:"column:title;" db:"title"`
	Archived  *bool   `json:"archived" gorm:"column:archived" db:"archived"`
	SummaryID *string `json:"summary_id,omitempty" gorm:"column:summary_id"`
	MindmapID *string `json:"mindmap_id,omitempty" gorm:"column:mindmap_id"`
}

func (api *api) UpdateNoteHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("note-id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var req NoteDataUpdateRequest

		if err := c.ShouldBind(&req); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var updateData entity.NoteDataUpdate
		updateData.Title = req.Title
		updateData.Archived = req.Archived

		if req.SummaryID != nil {
			summaryId, err := core.FromBase58(*req.SummaryID)
			if err != nil {
				common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
				return
			}
			updateData.SummaryID = new(int64)
			*updateData.SummaryID = int64(summaryId.GetLocalID())
		}

		if req.MindmapID != nil {
			mindmapId, err := core.FromBase58(*req.MindmapID)
			if err != nil {
				common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
				return
			}
			updateData.MindmapID = new(int64)
			*updateData.MindmapID = int64(mindmapId.GetLocalID())
		}

		requester := c.MustGet(core.KeyRequester).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateNote(ctx, int(uid.GetLocalID()), &updateData); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
