package api

import (
	"net/http"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) GetUnreadCountHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(core.Requester)
		ctx := core.ContextWithRequester(c.Request.Context(), requester)

		count, err := api.business.GetUnreadCount(ctx)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(map[string]interface{}{
			"unread_count": count,
		}))
	}
}
