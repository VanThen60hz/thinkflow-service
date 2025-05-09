package api

import (
	"net/http"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) MarkNotificationAsReadHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		notiId := c.Param("noti-id")
		if notiId == "" {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError("notification id is required"))
			return
		}

		err := api.business.MarkNotificationAsRead(c.Request.Context(), notiId)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(true, nil, nil))
	}
}
