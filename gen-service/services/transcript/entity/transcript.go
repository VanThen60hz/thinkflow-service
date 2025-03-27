package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Transcript struct {
	core.SQLModel
	Content string `json:"content" gorm:"column:content;" db:"content"`
}

func (Transcript) TableName() string { return "transcripts" }

func (t *Transcript) Mask() {
	t.SQLModel.Mask(common.MaskTypeTranscript)
}
