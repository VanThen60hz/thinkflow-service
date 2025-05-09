package entity

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/VanThen60hz/service-context/core"
)

type NotificationCreation struct {
	core.SQLModel
	NotiType       string         `json:"noti_type" gorm:"column:noti_type" db:"noti_type"`
	NotiSenderID   int64          `json:"noti_sender_id" gorm:"column:noti_sender_id" db:"noti_sender_id"`
	NotiReceivedID int64          `json:"noti_received_id" gorm:"column:noti_received_id" db:"noti_received_id"`
	NotiContent    string         `json:"noti_content" gorm:"column:noti_content" db:"noti_content"`
	NotiOptions    sql.NullString `json:"noti_options,omitempty" gorm:"column:noti_options;type:json" db:"noti_options"`
	IsRead         bool           `json:"is_read" gorm:"column:is_read;default:false" db:"is_read"`
}

func (NotificationCreation) TableName() string {
	return Notification{}.TableName()
}

func (n *NotificationCreation) PrepareForInsert() {
	n.SQLModel = core.NewSQLModel()
	if n.NotiOptions.Valid == false {
		n.NotiOptions = sql.NullString{String: "{}", Valid: true}
	}
	n.IsRead = false
}

func (n *NotificationCreation) Validate() error {
	n.NotiType = strings.TrimSpace(n.NotiType)
	n.NotiContent = strings.TrimSpace(n.NotiContent)

	if n.NotiType == "" {
		return errors.New("notification type is required")
	}

	if !ValidNotiTypes[n.NotiType] {
		return errors.New("invalid notification type")
	}

	if n.NotiContent == "" {
		return errors.New("notification content is required")
	}

	if n.NotiSenderID == 0 || n.NotiReceivedID == 0 {
		return errors.New("sender and receiver IDs are required")
	}

	return nil
}

// Optional: Update struct if needed
type NotificationUpdate struct {
	NotiContent *string         `json:"noti_content,omitempty" gorm:"column:noti_content" db:"noti_content"`
	NotiOptions *sql.NullString `json:"noti_options,omitempty" gorm:"column:noti_options;type:json" db:"noti_options"`
	IsRead      *bool           `json:"is_read,omitempty" gorm:"column:is_read" db:"is_read"`
	UpdatedAt   *time.Time      `json:"-" gorm:"column:updated_at" db:"updated_at"`
}

func (NotificationUpdate) TableName() string {
	return Notification{}.TableName()
}

func (n *NotificationUpdate) Validate() error {
	if n.NotiContent != nil {
		trimmed := strings.TrimSpace(*n.NotiContent)
		if trimmed == "" {
			return errors.New("notification content cannot be empty")
		}
		n.NotiContent = &trimmed
	}
	return nil
}

type Filter struct {
	NotiSenderID   *string `json:"noti_sender_id,omitempty" form:"noti_sender_id"`
	NotiReceivedID *string `json:"noti_received_id,omitempty" form:"noti_received_id"`
	NotiType       *string `json:"noti_type,omitempty" form:"noti_type"`
	IsRead         *bool   `json:"is_read,omitempty" form:"is_read"`
}
