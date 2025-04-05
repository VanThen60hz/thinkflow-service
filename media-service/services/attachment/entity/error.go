package entity

import "errors"

var (
	ErrAttachmentNotFound            = errors.New("attachment not found")
	ErrFileURLCannotBeBlank          = errors.New("attachment file URL cannot be blank")
	ErrCannotCreateAttachment        = errors.New("cannot create attachment")
	ErrCannotUpdateAttachment        = errors.New("cannot update attachment")
	ErrCannotDeleteAttachment        = errors.New("cannot delete attachment")
	ErrCannotListAttachments         = errors.New("cannot list attachments")
	ErrCannotGetAttachment           = errors.New("cannot get attachment details")
	ErrRequesterIsNotAttachmentOwner = errors.New("no permission, only attachment owner can do this")
	ErrNoteIDCannotBeBlank           = errors.New("note ID cannot be blank")
)
