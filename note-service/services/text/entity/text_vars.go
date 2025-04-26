package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

// TextDataCreation use for inserting data into database, we don't need all data fields
type TextDataCreation struct {
	core.SQLModel
	NoteID      int64          `json:"note_id" gorm:"column:note_id"`
	TextContent datatypes.JSON `json:"text_content" gorm:"column:text_content;type:json;" db:"text_content"`
	// UserId int `json:"-" gorm:"column:user_id" db:"user_id"`
}

func (TextDataCreation) TableName() string { return Text{}.TableName() }

func (text *TextDataCreation) Prepare(userId int) {
	text.SQLModel = core.NewSQLModel()
	// Text.UserId = userId
}

func (text *TextDataCreation) Mask() {
	text.SQLModel.Mask(common.MaskTypeText)
}

func (text *TextDataCreation) Validate() error {
	if len(text.TextContent) == 0 {
		return ErrTextContentCannotNotBlank
	}

	return nil
}

// TextDataUpdate contains only data fields can be used for updating
type TextDataUpdate struct {
	TextContent datatypes.JSON `json:"text_content" gorm:"column:text_content;type:json;" db:"text_content"`
	SummaryID   *int64         `json:"summary_id,omitempty" gorm:"column:summary_id"`
}

func (TextDataUpdate) TableName() string { return Text{}.TableName() }

func (text *TextDataUpdate) Validate() error {
	if len(text.TextContent) == 0 {
		return ErrTextContentCannotNotBlank
	}

	return nil
}

type Filter struct {
	NoteID *int64 `json:"note_id" form:"note_id"`
}
