package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Audio struct {
	core.SQLModel
	NoteID       int64                    `json:"-" gorm:"column:note_id"`
	FileURL      string                   `json:"file_url" gorm:"column:file_url"`
	Format       string                   `json:"format" gorm:"column:format"`
	TranscriptID *int64                   `json:"-" gorm:"column:transcript_id"`
	Transcript   *common.SimpleTranscript `json:"transcript,omitempty" gorm:"-" db:"-"`
	SummaryID    *int64                   `json:"-" gorm:"column:summary_id"`
	Summary      *common.SimpleSummary    `json:"summary,omitempty" gorm:"-" db:"-"`
}

func (Audio) TableName() string {
	return "audios"
}

func (au *Audio) Mask() {
	au.SQLModel.Mask(common.MaskTypeAudio)

	if t := au.Transcript; t != nil {
		t.Mask(common.MaskTypeTranscript)
	}

	if s := au.Summary; s != nil {
		s.Mask(common.MaskTypeSummary)
	}
}
