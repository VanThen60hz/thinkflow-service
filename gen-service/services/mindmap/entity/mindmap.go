package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type Mindmap struct {
	core.SQLModel
	MindmapData datatypes.JSON `json:"mindmap_data" gorm:"column:mindmap_data;" db:"mindmap_data"`
}

func (Mindmap) TableName() string { return "mindmaps" }

func (m *Mindmap) Mask() {
	m.SQLModel.Mask(common.MaskTypeMindmap)
}
