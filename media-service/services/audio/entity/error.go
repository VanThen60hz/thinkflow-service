package entity

import (
	"errors"
)

var (
	ErrUrlIsBlank            = errors.New("url cannot be blank")
	ErrUserIdNotValid        = errors.New("user id is not valid")
	ErrCannotGetPermission   = errors.New("cannot get permission")
	ErrCannotGetNoteByID     = errors.New("cannot get note by ID")
	ErrAudioNotFound         = errors.New("audio not found")
	ErrCannotCreateAudio     = errors.New("cannot create audio")
	ErrCannotUpdateAudio     = errors.New("cannot update audio")
	ErrCannotDeleteAudio     = errors.New("cannot delete audio")
	ErrCannotListAudio       = errors.New("cannot list audio")
	ErrCannotGetAudio        = errors.New("cannot get audio details")
	ErrRequesterIsNotOwner   = errors.New("no permission, only audio owner can do this")
	ErrCannotGetTranscript   = errors.New("cannot get transcript")
	ErrCannotGetSummary      = errors.New("cannot get summary")
	ErrRequesterCannotRead   = errors.New("requester cannot read permission")
	ErrRequesterCannotModify = errors.New("requester cannot modify permission")
)
