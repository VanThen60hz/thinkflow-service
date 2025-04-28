package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetAttachmentHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := core.FromBase58(c.Param("attachment-id"))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		attachment, err := api.business.GetAttachment(ctx, int64(id.GetLocalID()))
		if err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		attachment.Mask()
		c.JSON(http.StatusOK, core.ResponseData(attachment))
	}
}
