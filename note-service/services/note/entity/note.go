package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Note struct {
	core.SQLModel
	UserId     int                   `json:"-" gorm:"column:user_id" db:"user_id"`
	Title      string                `json:"title" gorm:"column:title;" db:"title"`
	Archived   bool                  `json:"archived" gorm:"column:archived;default:false" db:"archived"`
	User       *core.SimpleUser      `json:"user" gorm:"-" db:"-"`
	Permission string                `json:"permission" gorm:"-" db:"-"` // owner / read / write
	SummaryID  *int64                `json:"-" gorm:"column:summary_id"`
	Summary    *common.SimpleSummary `json:"summary,omitempty" gorm:"-" db:"-"`
	MindmapID  *int64                `json:"-" gorm:"column:mindmap_id"`
	Mindmap    *common.SimpleMindmap `json:"mindmap,omitempty" gorm:"-" db:"-"`
}

func (Note) TableName() string { return "notes" }

func (note *Note) Mask() {
	note.SQLModel.Mask(common.MaskTypeNote)

	if u := note.User; u != nil {
		u.Mask(common.MaskTypeUser)
	}

	if m := note.Mindmap; m != nil {
		m.Mask(common.MaskTypeMindmap)
	}
}

type NoteMember struct {
	*core.SimpleUser
	Role       string `json:"role"`       // owner / collaborator
	Permission string `json:"permission"` // owner / read / write
}
