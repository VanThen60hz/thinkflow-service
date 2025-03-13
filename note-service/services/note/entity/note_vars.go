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

func (t *NoteDataCreation) Prepare(userId int) {
	t.SQLModel = core.NewSQLModel()
	t.UserId = userId
}

func (t *NoteDataCreation) Mask() {
	t.SQLModel.Mask(common.MaskTypeNote)
}

func (t *NoteDataCreation) Validate() error {
	t.Title = strings.TrimSpace(t.Title)

	if err := checkTitle(t.Title); err != nil {
		return err
	}

	if t.UserId <= 0 {
		return ErrUserIdNotValid
	}

	return nil
}

// NoteDataUpdate contains only data fields can be used for updating
type NoteDataUpdate struct {
	Title *string `json:"title" gorm:"column:title;" db:"title"`
}

func (NoteDataUpdate) TableName() string { return Note{}.TableName() }

func (t *NoteDataUpdate) Validate() error {
	if title := t.Title; title != nil {
		s := strings.TrimSpace(*title)

		if err := checkTitle(s); err != nil {
			return err
		}

		t.Title = &s
	}

	return nil
}

type Filter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
}
