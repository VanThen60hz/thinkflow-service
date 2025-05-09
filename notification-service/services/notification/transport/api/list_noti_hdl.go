package api

import (
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/notification/dto"
	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

func (api *api) ListNotificationsHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		type reqParam struct {
			entity.Filter
			core.Paging
		}

		var rp reqParam

		if err := c.ShouldBind(&rp); err != nil {
			core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		requester := c.MustGet(common.RequesterKey).(core.Requester)
		requesterSubject := requester.GetSubject()

		rp.Paging.Process()
		rp.NotiReceivedID = &requesterSubject

		notis, err := api.business.ListNotifications(c.Request.Context(), &rp.Filter, &rp.Paging)
		if err != nil {
			core.WriteErrorResponse(c, err)
			return
		}

		for i := range notis {
			notis[i].Mask()
		}

		response := dto.ListNotificationResponse{
			Data:   make([]dto.NotificationResponse, len(notis)),
			Paging: rp.Paging,
		}

		for i := range notis {
			response.Data[i] = *dto.NewNotificationResponse(&notis[i])
		}

		c.JSON(http.StatusOK, core.SuccessResponse(response.Data, response.Paging, rp.Filter))
	}
}
