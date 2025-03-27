package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type MindmapDataCreation struct {
	core.SQLModel
	MindmapData datatypes.JSON `json:"mindmap_data" gorm:"column:mindmap_data;" db:"mindmap_data"`
}

func (MindmapDataCreation) TableName() string { return Mindmap{}.TableName() }

func (m *MindmapDataCreation) Prepare(userId int) {
	m.SQLModel = core.NewSQLModel()
}

func (m *MindmapDataCreation) Mask() {
	m.SQLModel.Mask(common.MaskTypeMindmap)
}

func (m *MindmapDataCreation) Validate() error {
	if len(m.MindmapData) == 0 {
		return ErrMindmapDataNotValid
	}

	return nil
}

type MindmapDataUpdate struct {
	MindmapData datatypes.JSON `json:"mindmap_data" gorm:"column:mindmap_data;" db:"mindmap_data"`
}

func (MindmapDataUpdate) TableName() string { return Mindmap{}.TableName() }

func (m *MindmapDataUpdate) Validate() error {
	if len(m.MindmapData) == 0 {
		return ErrMindmapDataNotValid
	}

	return nil
}

type MindmapFilter struct {
	UserId *string `json:"user_id,omitempty" form:"user_id"`
}
