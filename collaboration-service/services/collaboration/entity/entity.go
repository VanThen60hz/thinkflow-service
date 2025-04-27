package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type PermissionType string

const (
	PermissionRead  PermissionType = "read"
	PermissionWrite PermissionType = "write"
)

func (p PermissionType) IsValid() bool {
	switch p {
	case PermissionRead, PermissionWrite:
		return true
	default:
		return false
	}
}

type Collaboration struct {
	core.SQLModel
	NoteId     int            `json:"note_id" gorm:"column:note_id" db:"note_id"`
	UserId     int            `json:"user_id" gorm:"column:user_id" db:"user_id"`
	Permission PermissionType `json:"permission" gorm:"column:permission" db:"permission"`

	User *core.SimpleUser `json:"user,omitempty" gorm:"-" db:"-"`
}

func (Collaboration) TableName() string { return "collaborations" }

func (c *Collaboration) Mask() {
	c.SQLModel.Mask(common.MaskTypeCollaboration)

	if c.User != nil {
		c.User.Mask(common.MaskTypeUser)
	}
}
