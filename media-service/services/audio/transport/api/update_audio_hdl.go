package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type UpdateAudioRequest struct {
	TranscriptID *string `json:"transcript_id,omitempty" gorm:"column:transcript_id"`
	SummaryID    *string `json:"summary_id,omitempty" gorm:"column:summary_id"`
}

func (api *api) UpdateAudioHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var req UpdateAudioRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		updateData := entity.AudioDataUpdate{}

		if req.TranscriptID != nil {
			transcriptId, err := core.FromBase58(*req.TranscriptID)
			if err != nil {
				core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
				return
			}
			updateData.TranscriptID = new(int64)
			*updateData.TranscriptID = int64(transcriptId.GetLocalID())
		}

		if req.SummaryID != nil {
			summaryId, err := core.FromBase58(*req.SummaryID)
			if err != nil {
				core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
				return
			}
			updateData.SummaryID = new(int64)
			*updateData.SummaryID = int64(summaryId.GetLocalID())
		}

		audioId, err := core.FromBase58(c.Param("audio-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateAudio(ctx, int(audioId.GetLocalID()), &updateData); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
