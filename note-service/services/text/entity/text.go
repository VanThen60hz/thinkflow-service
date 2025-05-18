package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type Text struct {
	core.SQLModel
	NoteID      int64                 `json:"note_id" gorm:"column:note_id;" db:"note_id"`
	TextContent datatypes.JSON        `json:"text_content" gorm:"column:text_content;type:json;" db:"text_content"`
	TextString  string                `json:"text_string" gorm:"column:text_string;type:text;" db:"text_string"`
	SummaryID   *int64                `json:"-" gorm:"column:summary_id"`
	Summary     *common.SimpleSummary `json:"summary,omitempty" gorm:"-" db:"-"`
}

func (Text) TableName() string { return "texts" }

func (text *Text) Mask() {
	text.SQLModel.Mask(common.MaskTypeText)

	if s := text.Summary; s != nil {
		s.Mask(common.MaskTypeSummary)
	}
}
