package entity

import (
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

// NoteDataCreation use for inserting data into database, we don't need all data fields
type NoteDataCreation struct {
	core.SQLModel
	Title string `json:"title" gorm:"column:title;" db:"title"`
	// Do not allow client set these fields
	UserId int `json:"-" gorm:"column:user_id" db:"user_id"`
}

func (NoteDataCreation) TableName() string { return Note{}.TableName() }

func (note *NoteDataCreation) Prepare(userId int) {
	note.SQLModel = core.NewSQLModel()
	note.UserId = userId
}

func (note *NoteDataCreation) Mask() {
	note.SQLModel.Mask(common.MaskTypeNote)
}

func (note *NoteDataCreation) Validate() error {
	note.Title = strings.TrimSpace(note.Title)

	if err := checkTitle(note.Title); err != nil {
		return err
	}

	if note.UserId <= 0 {
		return ErrUserIdNotValid
	}

	return nil
}

// NoteDataUpdate contains only data fields can be used for updating
type NoteDataUpdate struct {
	Title     *string `json:"title" gorm:"column:title;" db:"title"`
	Archived  *bool   `json:"archived" gorm:"column:archived" db:"archived"`
	SummaryID *int64  `json:"summary_id,omitempty" gorm:"column:summary_id"`
	MindmapID *int64  `json:"mindmap_id,omitempty" gorm:"column:mindmap_id"`
}

func (NoteDataUpdate) TableName() string { return Note{}.TableName() }

func (note *NoteDataUpdate) Validate() error {
	if title := note.Title; title != nil {
		s := strings.TrimSpace(*title)

		if err := checkTitle(s); err != nil {
			return err
		}

		note.Title = &s
	}

	return nil
}

type Filter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
	Title  *string `json:"title,omitempty" form:"title"`
}
