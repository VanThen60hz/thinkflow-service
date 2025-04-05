package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type AttachmentCreation struct {
	core.SQLModel
	NoteID    int64  `json:"note_id" gorm:"column:note_id"`
	FileURL   string `json:"file_url" gorm:"column:file_url"`
	FileName  string `json:"file_name" gorm:"column:file_name"`
	Extension string `json:"extension" gorm:"column:extension"`
	SizeBytes int64  `json:"size_bytes" gorm:"column:size_bytes"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name"`
}

func (AttachmentCreation) TableName() string { return Attachment{}.TableName() }

func (att *AttachmentCreation) Prepare(userId int) {
	att.SQLModel = core.NewSQLModel()
}

func (att *AttachmentCreation) Mask() {
	att.SQLModel.Mask(common.MaskTypeAttachment)
}

func (att *AttachmentCreation) Validate() error {
	if att.NoteID == 0 {
		return ErrNoteIDCannotBeBlank
	}
	if att.FileURL == "" {
		return ErrFileURLCannotBeBlank
	}
	return nil
}
