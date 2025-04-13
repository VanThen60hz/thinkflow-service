package entity

import (
	"time"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type NoteShareLink struct {
	core.SQLModel
	NoteID     int64      `json:"note_id" gorm:"column:note_id;" db:"note_id"`
	Permission string     `json:"permission" gorm:"column:permission;type:enum('read','write');" db:"permission"`
	Token      string     `json:"token" gorm:"column:token;unique;" db:"token"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty" gorm:"column:expires_at;default:null;" db:"expires_at"`
	CreatedBy  int        `json:"created_by" gorm:"column:created_by;" db:"created_by"`
}

func (NoteShareLink) TableName() string { return "note_share_links" }

func (link *NoteShareLink) Mask() {
	link.SQLModel.Mask(common.MaskTypeShareLink)
}
