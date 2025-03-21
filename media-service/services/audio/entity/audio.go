package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Audio struct {
	core.SQLModel
	NoteID       int64  `json:"-" gorm:"column:note_id"`
	FileURL      string `json:"file_url" gorm:"column:file_url"`
	Format       string `json:"format" gorm:"column:format"`
	TranscriptID *int64 `json:"transcript_id,omitempty" gorm:"column:transcript_id"`
	SummaryID    *int64 `json:"summary_id,omitempty" gorm:"column:summary_id"`
	MindmapID    *int64 `json:"mindmap_id,omitempty" gorm:"column:mindmap_id"`
}

func (Audio) TableName() string {
	return "audios"
}

func (au *Audio) Mask() {
	au.SQLModel.Mask(common.MaskTypeImage)
}
