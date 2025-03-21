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

func (img *ImageDataCreation) Prepare() {
	img.SQLModel = core.NewSQLModel()
}

func (img *ImageDataCreation) Mask() {
	img.SQLModel.Mask(common.MaskTypeImage)
}

func (img *ImageDataCreation) Validate() error {
	img.Url = strings.TrimSpace(img.Url)

	if img.Url == "" {
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

func (img *ImageDataUpdate) Validate() error {
	if url := img.Url; url != nil {
		s := strings.TrimSpace(*url)

		if s == "" {
			return ErrUrlIsBlank
		}

		img.Url = &s
	}

	return nil
}

type Filter struct {
	Extension *string `json:"extension" form:"extension" db:"extension"`
	Folder    *string `json:"folder" form:"folder" db:"folder"`
	CloudName *string `json:"cloud_name" form:"cloud_name" db:"cloud_name"`
}
