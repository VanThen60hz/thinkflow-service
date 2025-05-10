package api

import (
	"database/sql"
	"net/http"

	"thinkflow-service/common"
	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type createNotificationRequest struct {
	NotiType       string         `json:"noti_type" binding:"required"`
	NotiSenderID   string         `json:"noti_sender_id" binding:"required"`
	NotiReceivedID string         `json:"noti_received_id" binding:"required"`
	NotiContent    string         `json:"noti_content" binding:"required"`
	NotiOptions    sql.NullString `json:"noti_options"`
}

func (hdl *api) CreateNotification(c *gin.Context) {
	var req createNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	notiSenderUid, err := core.FromBase58(req.NotiSenderID)
	if err != nil {
		core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	notiRecievedUid, err := core.FromBase58(req.NotiReceivedID)
	if err != nil {
		core.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	noti := &entity.NotificationCreation{
		NotiType:       req.NotiType,
		NotiSenderID:   int64(notiSenderUid.GetLocalID()),
		NotiReceivedID: int64(notiRecievedUid.GetLocalID()),
		NotiContent:    req.NotiContent,
		NotiOptions:    req.NotiOptions,
	}

	requester := c.MustGet(common.RequesterKey).(core.Requester)
	ctx := core.ContextWithRequester(c.Request.Context(), requester)

	if err := hdl.business.CreateNotification(ctx, noti); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, core.ResponseData(map[string]interface{}{
		"data": noti,
	}))
}
