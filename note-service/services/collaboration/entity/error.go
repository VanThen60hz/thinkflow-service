package entity

import "errors"

var (
	ErrInvalidPermission     = errors.New("permission must be either 'read' or 'write'")
	ErrUserIdInvalid         = errors.New("user id is not valid")
	ErrNoteIdInvalid         = errors.New("note id is not valid")
	ErrNoteIdCannotBeBlank   = errors.New("note id cannot be blank")
	ErrUserIdCannotBeBlank   = errors.New("user id cannot be blank")
	ErrCollaborationExists   = errors.New("collaboration already exists")
	ErrCollaborationNotFound = errors.New("collaboration not found")
	ErrCannotCreateCollab    = errors.New("cannot create collaboration")
	ErrCannotUpdateCollab    = errors.New("cannot update collaboration")
	ErrCannotDeleteCollab    = errors.New("cannot delete collaboration")
	ErrCannotListCollab      = errors.New("cannot list collaborations")
	ErrRequesterNotNoteOwner = errors.New("no permission, only note owner can share")
	ErrRequesterCannotModify = errors.New("no permission to modify this collaboration")
)
