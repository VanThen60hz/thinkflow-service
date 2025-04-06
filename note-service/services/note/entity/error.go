package entity

import "errors"

var (
	ErrTitleIsBlank             = errors.New("title cannot be blank")
	ErrUserIdNotValid           = errors.New("user id is not valid")
	ErrNoteNotFound             = errors.New("note not found")
	ErrCannotCreateNote         = errors.New("cannot create note")
	ErrCannotUpdateNote         = errors.New("cannot update note")
	ErrCannotDeleteNote         = errors.New("cannot update note")
	ErrCannotListNote           = errors.New("cannot list notes")
	ErrCannotListArchivedNote   = errors.New("cannot list archived notes")
	ErrCannotGetNote            = errors.New("cannot get note details")
	ErrRequesterIsNotOwner      = errors.New("no permission, only note owner can do this")
	ErrCannotArchiveNote        = errors.New("cannot archive note")
	ErrCannotUnarchiveNote      = errors.New("cannot unarchive note")
	ErrCannotUpdateArchivedNote = errors.New("cannot update archived note")
)
