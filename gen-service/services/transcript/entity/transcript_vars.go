package entity

import (
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

// TranscriptDataCreation dùng để chèn dữ liệu vào database, không cần tất cả các field
type TranscriptDataCreation struct {
	core.SQLModel
	Content string `json:"content" gorm:"column:content;" db:"content"`
	// Không cho phép client set trường này
	// UserId int `json:"-" gorm:"column:user_id" db:"user_id"`
}

func (TranscriptDataCreation) TableName() string { return Transcript{}.TableName() }

func (t *TranscriptDataCreation) Prepare(userId int) {
	t.SQLModel = core.NewSQLModel()
	// t.UserId = userId
}

func (t *TranscriptDataCreation) Mask() {
	t.SQLModel.Mask(common.MaskTypeTranscript)
}

func (t *TranscriptDataCreation) Validate() error {
	t.Content = strings.TrimSpace(t.Content)

	if len(t.Content) == 0 {
		return ErrContentNotValid
	}

	// if t.UserId <= 0 {
	// 	return ErrUserIdNotValid
	// }

	return nil
}

type TranscriptDataUpdate struct {
	Content *string `json:"content" gorm:"column:content;" db:"content"`
}

func (TranscriptDataUpdate) TableName() string { return Transcript{}.TableName() }

func (t *TranscriptDataUpdate) Validate() error {
	if content := t.Content; content != nil {
		s := strings.TrimSpace(*content)

		if len(s) == 0 {
			return ErrContentNotValid
		}

		t.Content = &s
	}

	return nil
}

type TranscriptFilter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
}
