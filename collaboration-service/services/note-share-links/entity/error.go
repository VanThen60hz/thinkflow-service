package entity

import "errors"

var (
	ErrInvalidNoteID             = errors.New("note_id is not valid")
	ErrInvalidPermission         = errors.New("permission must be 'read' or 'write'")
	ErrTokenCannotBeBlank        = errors.New("token cannot be blank")
	ErrNoteShareLinkNotFound     = errors.New("note share link not found")
	ErrCannotCreateShareLink     = errors.New("cannot create note share link")
	ErrCannotUpdateShareLink     = errors.New("cannot update note share link")
	ErrCannotDeleteShareLink     = errors.New("cannot delete note share link")
	ErrCannotListShareLinks      = errors.New("cannot list note share links")
	ErrCannotGetShareLink        = errors.New("cannot get note share link details")
	ErrRequesterIsNotOwner       = errors.New("no permission, only link creator can do this")
	ErrRequesterCannotModifyLink = errors.New("no permission to modify this link")
)
