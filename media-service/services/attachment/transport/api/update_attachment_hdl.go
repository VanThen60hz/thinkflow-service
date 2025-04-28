package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) UpdateAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("attachment-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		var data entity.Attachment
		if err := c.ShouldBindJSON(&data); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.UpdateAttachment(ctx, int64(id.GetLocalID()), &data); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
