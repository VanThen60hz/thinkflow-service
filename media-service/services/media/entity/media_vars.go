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
	ErrMediaNotFound       = errors.New("media not found")
	ErrCannotCreateMedia   = errors.New("cannot create media")
	ErrCannotUpdateMedia   = errors.New("cannot update media")
	ErrCannotDeleteMedia   = errors.New("cannot delete media")
	ErrCannotListMedia     = errors.New("cannot list media")
	ErrCannotGetMedia      = errors.New("cannot get media details")
	ErrRequesterIsNotOwner = errors.New("no permission, only media owner can do this")
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
	t.SQLModel.Mask(common.MaskTypeMedia)
}

func (t *ImageDataCreation) Validate() error {
	t.Url = strings.TrimSpace(t.Url)

	if t.Url == "" {
		return ErrUrlIsBlank
	}

	return nil
}

type AudioDataCreation struct {
	core.SQLModel
	Url        string `json:"url" gorm:"column:url;" db:"url"`
	Format     string `json:"format" gorm:"column:format;" db:"format"`
	Duration   int64  `json:"duration" gorm:"column:duration;" db:"duration"`
	UploadedAt string `json:"uploaded_at" gorm:"column:uploaded_at;" db:"uploaded_at"`
	UserId     int    `json:"-" gorm:"column:user_id" db:"user_id"`
}

func (AudioDataCreation) TableName() string { return Audio{}.TableName() }

func (t *AudioDataCreation) Prepare(userId int) {
	t.SQLModel = core.NewSQLModel()
	t.UserId = userId
}

func (t *AudioDataCreation) Mask() {
	t.SQLModel.Mask(common.MaskTypeMedia)
}

func (t *AudioDataCreation) Validate() error {
	t.Url = strings.TrimSpace(t.Url)

	if t.Url == "" {
		return ErrUrlIsBlank
	}

	if t.UserId <= 0 {
		return ErrUserIdNotValid
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

type AudioDataUpdate struct {
	Url        *string `json:"url" gorm:"column:url;" db:"url"`
	Format     *string `json:"format" gorm:"column:format;" db:"format"`
	Duration   *int64  `json:"duration" gorm:"column:duration;" db:"duration"`
	UploadedAt *string `json:"uploaded_at" gorm:"column:uploaded_at;" db:"uploaded_at"`
}

func (AudioDataUpdate) TableName() string { return Audio{}.TableName() }

func (t *AudioDataUpdate) Validate() error {
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
	Type string `json:"type,omitempty" form:"type"` // "image" or "audio"
}
