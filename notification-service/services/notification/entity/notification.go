package entity

import (
	"database/sql"
	"errors"
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Notification struct {
	core.SQLModel
	NotiType       string         `json:"noti_type" gorm:"column:noti_type"`
	NotiSenderID   int64          `json:"noti_sender_id" gorm:"column:noti_sender_id"`
	NotiReceivedID int64          `json:"noti_received_id" gorm:"column:noti_received_id"`
	NotiContent    string         `json:"noti_content" gorm:"column:noti_content"`
	NotiOptions    sql.NullString `json:"noti_options" gorm:"column:noti_options;type:json"`
	IsRead         bool           `json:"is_read" gorm:"column:is_read;default:false"`

	Sender   *core.SimpleUser `json:"sender" gorm:"-"`   // virtual field
	Receiver *core.SimpleUser `json:"receiver" gorm:"-"` // virtual field
}

func (Notification) TableName() string {
	return "notifications"
}

func (n *Notification) Mask() {
	n.SQLModel.Mask(common.MaskTypeNotification)

	if n.Sender != nil {
		n.Sender.Mask(common.MaskTypeUser)
	}
	if n.Receiver != nil {
		n.Receiver.Mask(common.MaskTypeUser)
	}
}

func (n *Notification) Validate() error {
	n.NotiType = strings.TrimSpace(n.NotiType)
	n.NotiContent = strings.TrimSpace(n.NotiContent)

	if n.NotiType == "" {
		return errors.New("notification type is required")
	}

	if n.NotiContent == "" {
		return errors.New("notification content is required")
	}

	return nil
}

var ValidNotiTypes = map[string]bool{
	"NOTE_CREATED":         true,
	"TRANSCRIPT_GENERATED": true,
	"SUMMARY_GENERATED":    true,
	"MINDMAP_GENERATED":    true,
	"AUDIO_PROCESSED":      true,
	"TEXT_PROCESSED":       true,
	"REMINDER":             true,
	"COLLAB_INVITE":        true,
	"COMMENT":              true,
}
