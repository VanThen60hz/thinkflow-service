package entity

import (
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type SummaryDataCreation struct {
	core.SQLModel
	SummaryText string `json:"summary_text" gorm:"column:summary_text;" db:"summary_text"`
	// UserId      int    `json:"-" gorm:"column:user_id" db:"user_id"`
}

func (SummaryDataCreation) TableName() string { return Summary{}.TableName() }

func (s *SummaryDataCreation) Prepare(userId int) {
	s.SQLModel = core.NewSQLModel()
	// s.UserId = userId
}

func (s *SummaryDataCreation) Mask() {
	s.SQLModel.Mask(common.MaskTypeSummary)
}

func (s *SummaryDataCreation) Validate() error {
	s.SummaryText = strings.TrimSpace(s.SummaryText)

	if len(s.SummaryText) == 0 {
		return ErrContentNotValid
	}

	// if s.UserId <= 0 {
	// 	return ErrUserIdNotValid
	// }

	return nil
}

type SummaryDataUpdate struct {
	SummaryText *string `json:"summary_text" gorm:"column:summary_text;" db:"summary_text"`
}

func (SummaryDataUpdate) TableName() string { return Summary{}.TableName() }

func (s *SummaryDataUpdate) Validate() error {
	if summaryText := s.SummaryText; summaryText != nil {
		text := strings.TrimSpace(*summaryText)

		if len(text) == 0 {
			return ErrContentNotValid
		}

		s.SummaryText = &text
	}

	return nil
}

type SummaryFilter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
}
