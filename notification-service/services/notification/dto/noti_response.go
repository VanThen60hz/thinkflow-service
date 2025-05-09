package dto

import (
	"time"

	"thinkflow-service/services/notification/entity"

	"github.com/VanThen60hz/service-context/core"
)

type NotificationResponse struct {
	Id             string     `json:"id"`
	NotiType       string     `json:"noti_type"`
	NotiSenderID   int64      `json:"noti_sender_id"`
	NotiReceivedID int64      `json:"noti_received_id"`
	NotiContent    string     `json:"noti_content"`
	NotiOptions    *string    `json:"noti_options,omitempty"` // chuyển từ sql.NullString
	IsRead         bool       `json:"is_read"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

func NewNotificationResponse(noti *entity.Notification) *NotificationResponse {
	var options *string
	if noti.NotiOptions.Valid {
		options = &noti.NotiOptions.String
	}

	return &NotificationResponse{
		Id:             noti.FakeId.String(),
		NotiType:       noti.NotiType,
		NotiSenderID:   noti.NotiSenderID,
		NotiReceivedID: noti.NotiReceivedID,
		NotiContent:    noti.NotiContent,
		NotiOptions:    options,
		IsRead:         noti.IsRead,
		CreatedAt:      noti.CreatedAt,
		UpdatedAt:      noti.UpdatedAt,
	}
}

type ListNotificationResponse struct {
	Data   []NotificationResponse `json:"data"`
	Paging core.Paging            `json:"paging"`
}
