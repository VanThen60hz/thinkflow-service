package entity

import (
	"errors"
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

var (
	ErrUrlIsBlank          = errors.New("url cannot be blank")
	ErrUserIdNotValid      = errors.New("user id is not valid")
	ErrImageNotFound       = errors.New("image not found")
	ErrCannotCreateImage   = errors.New("cannot create image")
	ErrCannotUpdateImage   = errors.New("cannot update image")
	ErrCannotDeleteImage   = errors.New("cannot delete image")
	ErrCannotListImage     = errors.New("cannot list image")
	ErrCannotGetImage      = errors.New("cannot get image details")
	ErrRequesterIsNotOwner = errors.New("no permission, only image owner can do this")
)

type ImageDataCreation struct {
	core.SQLModel
	Url       string `json:"url" gorm:"column:url;" db:"url"`
	Width     int64  `json:"width" gorm:"column:width;" db:"width"`
	Height    int64  `json:"height" gorm:"column:height;" db:"height"`
	Extension string `json:"extension" gorm:"column:extension;" db:"extension"`
	Folder    string `json:"folder" gorm:"column:folder;" db:"folder"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name;" db:"cloud_name"`
}

func (ImageDataCreation) TableName() string { return Image{}.TableName() }

func (t *ImageDataCreation) Prepare() {
	t.SQLModel = core.NewSQLModel()
}

func (t *ImageDataCreation) Mask() {
	t.SQLModel.Mask(common.MaskTypeImage)
}

func (t *ImageDataCreation) Validate() error {
	t.Url = strings.TrimSpace(t.Url)

	if t.Url == "" {
		return ErrUrlIsBlank
	}

	return nil
}

type ImageDataUpdate struct {
	Url       *string `json:"url" gorm:"column:url;" db:"url"`
	Width     *int64  `json:"width" gorm:"column:width;" db:"width"`
	Height    *int64  `json:"height" gorm:"column:height;" db:"height"`
	Extension *string `json:"extension" gorm:"column:extension;" db:"extension"`
	Folder    *string `json:"folder" gorm:"column:folder;" db:"folder"`
	CloudName *string `json:"cloud_name" gorm:"column:cloud_name;" db:"cloud_name"`
}

func (ImageDataUpdate) TableName() string { return Image{}.TableName() }

func (t *ImageDataUpdate) Validate() error {
	if url := t.Url; url != nil {
		s := strings.TrimSpace(*url)

		if s == "" {
			return ErrUrlIsBlank
		}

		t.Url = &s
	}

	return nil
}

type Filter struct {
	Type string `json:"type,omitempty" form:"type"`
}
