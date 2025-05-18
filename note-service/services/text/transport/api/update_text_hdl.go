package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/text/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type TextUpdateRequest struct {
	TextContent datatypes.JSON `json:"text_content" gorm:"column:text_content;type:json;" db:"text_content"`
	TextString  string         `json:"text_string" gorm:"column:text_string;type:text;" db:"text_string"`
	SummaryID   *string        `json:"summary_id,omitempty" gorm:"column:summary_id"`
}

func (api *api) UpdateTextHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := core.FromBase58(c.Param("text-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var req TextUpdateRequest

		if err := c.ShouldBind(&req); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var updateData entity.TextDataUpdate

		if req.TextContent != nil {
			updateData.TextContent = req.TextContent
		}

		if req.TextString != "" {
			updateData.TextString = req.TextString
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

		// Set requester to context
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateText(ctx, int(uid.GetLocalID()), &updateData); err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
