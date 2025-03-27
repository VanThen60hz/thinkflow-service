package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Summary struct {
	core.SQLModel
	SummaryText string `json:"summary_text" gorm:"column:summary_text;" db:"summary_text"`
}

func (Summary) TableName() string { return "summaries" }

func (s *Summary) Mask() {
	s.SQLModel.Mask(common.MaskTypeSummary)
}
