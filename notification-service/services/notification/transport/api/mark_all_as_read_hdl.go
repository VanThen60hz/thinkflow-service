package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) MarkAllNotificationsAsReadHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		if err := api.business.MarkAllNotificationsAsRead(ctx); err != nil {
			core.WriteErrorResponse(c, core.ErrInternalServerError.WithError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}
