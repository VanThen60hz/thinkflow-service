package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Attachment struct {
	core.SQLModel
	NoteID    int64  `json:"-" gorm:"column:note_id"`
	FileURL   string `json:"file_url" gorm:"column:file_url"`
	FileName  string `json:"file_name" gorm:"column:file_name"`
	Extension string `json:"extension" gorm:"column:extension"`
	SizeBytes int64  `json:"size_bytes" gorm:"column:size_bytes"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name"`
}

func (Attachment) TableName() string {
	return "attachments"
}

func (a *Attachment) Mask() {
	a.SQLModel.Mask(common.MaskTypeAttachment)
}
