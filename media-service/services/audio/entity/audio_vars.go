package entity

import (
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type AudioDataCreation struct {
	core.SQLModel
	NoteID  int64  `json:"note_id" gorm:"column:note_id"`
	Name    string `json:"name" gorm:"column:name"`
	FileURL string `json:"file_url" gorm:"column:file_url"`
	Format  string `json:"format" gorm:"column:format"`
}

func (AudioDataCreation) TableName() string { return Audio{}.TableName() }

func (au *AudioDataCreation) Prepare() {
	au.SQLModel = core.NewSQLModel()
}

func (au *AudioDataCreation) Mask() {
	au.SQLModel.Mask(common.MaskTypeAudio)
}

func (au *AudioDataCreation) Validate() error {
	au.FileURL = strings.TrimSpace(au.FileURL)

	if au.FileURL == "" {
		return ErrUrlIsBlank
	}

	return nil
}

type AudioDataUpdate struct {
	core.SQLModel
	NoteID       *int64  `json:"note_id" gorm:"column:note_id"`
	FileURL      *string `json:"file_url" gorm:"column:file_url"`
	TranscriptID *int64  `json:"transcript_id,omitempty" gorm:"column:transcript_id"`
	SummaryID    *int64  `json:"summary_id,omitempty" gorm:"column:summary_id"`
}

func (AudioDataUpdate) TableName() string { return Audio{}.TableName() }

func (au *AudioDataUpdate) Validate() error {
	if url := au.FileURL; url != nil {
		s := strings.TrimSpace(*url)

		if s == "" {
			return ErrUrlIsBlank
		}

		au.FileURL = &s
	}

	return nil
}

type Filter struct {
	NoteID *int64 `json:"note_id,omitempty" form:"note_id"`
}
