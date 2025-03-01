package entity

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type Image struct {
	core.SQLModel
	Url       string `json:"url" gorm:"column:url;" db:"url"`
	Width     int64  `json:"width" gorm:"column:width;" db:"width"`
	Height    int64  `json:"height" gorm:"column:height;" db:"height"`
	Extension string `json:"extension" gorm:"column:extension;" db:"extension"`
	Folder    string `json:"folder" gorm:"column:folder;" db:"folder"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name;" db:"cloud_name"`
}

func (Image) TableName() string { return "images" }

func (t *Image) Mask() {
	t.SQLModel.Mask(common.MaskTypeMedia)
}

type Audio struct {
	core.SQLModel
	UserId     int              `json:"-" gorm:"column:user_id" db:"user_id"`
	Url        string           `json:"url" gorm:"column:url;" db:"url"`
	Format     string           `json:"format" gorm:"column:format;" db:"format"`
	Duration   int64            `json:"duration" gorm:"column:duration;" db:"duration"`
	UploadedAt string           `json:"uploaded_at" gorm:"column:uploaded_at;" db:"uploaded_at"`
	User       *core.SimpleUser `json:"user" gorm:"-" db:"-"`
}

func (Audio) TableName() string { return "audio_files" }

func (t *Audio) Mask() {
	t.SQLModel.Mask(common.MaskTypeMedia)

	if u := t.User; u != nil {
		u.Mask(common.MaskTypeUser)
	}
}
