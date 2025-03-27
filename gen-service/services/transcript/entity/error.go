package entity

import "errors"

var (
	ErrContentNotValid        = errors.New("content cannot be blank")
	ErrUserIdNotValid         = errors.New("user id is not valid")
	ErrTranscriptNotFound     = errors.New("transcript not found")
	ErrCannotCreateTranscript = errors.New("cannot create transcript")
	ErrCannotUpdateTranscript = errors.New("cannot update transcript")
	ErrCannotDeleteTranscript = errors.New("cannot delete transcript")
	ErrCannotListTranscript   = errors.New("cannot list transcripts")
	ErrCannotGetTranscript    = errors.New("cannot get transcript details")
	ErrRequesterIsNotOwner    = errors.New("no permission, only transcript owner can do this")
)
