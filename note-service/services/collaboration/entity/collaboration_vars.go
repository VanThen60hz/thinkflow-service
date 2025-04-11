package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type CollaborationCreation struct {
	core.SQLModel
	NoteId     int            `json:"note_id" gorm:"column:note_id"`
	UserId     int            `json:"user_id" gorm:"column:user_id"`
	Permission PermissionType `json:"permission" gorm:"column:permission"`
}

func (CollaborationCreation) TableName() string { return Collaboration{}.TableName() }

func (c *CollaborationCreation) Prepare() {
	c.SQLModel = core.NewSQLModel()
}

func (c *CollaborationCreation) Mask() {
	c.SQLModel.Mask(common.MaskTypeCollaboration)
}

func (c *CollaborationCreation) Validate() error {
	if c.NoteId == 0 {
		return ErrNoteIdCannotBeBlank
	}
	if c.UserId == 0 {
		return ErrUserIdCannotBeBlank
	}
	if !c.Permission.IsValid() {
		return ErrInvalidPermission
	}
	return nil
}

type CollaborationUpdate struct {
	Permission PermissionType `json:"permission" gorm:"column:permission"`
}

func (CollaborationUpdate) TableName() string { return Collaboration{}.TableName() }

func (c *CollaborationUpdate) Validate() error {
	if !c.Permission.IsValid() {
		return ErrInvalidPermission
	}
	return nil
}

type CollaborationFilter struct {
	NoteId *int `json:"note_id" form:"note_id"`
	UserId *int `json:"user_id" form:"user_id"`
}
