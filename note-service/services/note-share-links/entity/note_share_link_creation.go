package entity

import (
	"time"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type NoteShareLinkCreation struct {
	core.SQLModel
	NoteID     int64      `json:"note_id" gorm:"column:note_id"`
	Permission string     `json:"permission" gorm:"column:permission;type:enum('read','write')"`
	Token      string     `json:"token" gorm:"column:token"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty" gorm:"column:expires_at;default:null"`
	CreatedBy  int        `json:"created_by" gorm:"column:created_by"`
}

func (NoteShareLinkCreation) TableName() string { return NoteShareLink{}.TableName() }

func (link *NoteShareLinkCreation) Prepare(userId int) {
	link.SQLModel = core.NewSQLModel()
	link.CreatedBy = userId
}

func (link *NoteShareLinkCreation) Mask() {
	link.SQLModel.Mask(common.MaskTypeShareLink)
}

func (link *NoteShareLinkCreation) Validate() error {
	if link.NoteID <= 0 {
		return ErrInvalidNoteID
	}
	if link.Permission != "read" && link.Permission != "write" {
		return ErrInvalidPermission
	}
	if link.Token == "" {
		return ErrTokenCannotBeBlank
	}
	return nil
}
