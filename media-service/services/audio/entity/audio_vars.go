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
	ErrAudioNotFound       = errors.New("audio not found")
	ErrCannotCreateAudio   = errors.New("cannot create audio")
	ErrCannotUpdateAudio   = errors.New("cannot update audio")
	ErrCannotDeleteAudio   = errors.New("cannot delete audio")
	ErrCannotListAudio     = errors.New("cannot list audio")
	ErrCannotGetAudio      = errors.New("cannot get audio details")
	ErrRequesterIsNotOwner = errors.New("no permission, only audio owner can do this")
)

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
	t.SQLModel.Mask(common.MaskTypeAudio)
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
	Type string `json:"type,omitempty" form:"type"`
}
