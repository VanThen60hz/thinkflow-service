package entity

import "errors"
import "github.com/VanThen60hz/service-context/core"

var (
	ErrTextContentCannotNotBlank = core.ErrBadRequest.WithError("text content cannot be blank")
	ErrTextStringCannotNotBlank  = core.ErrBadRequest.WithError("text string cannot be blank")
	ErrUserIdNotValid            = errors.New("user id is not valid")
	ErrTextNotFound              = errors.New("text not found")
	ErrCannotCreateText          = errors.New("cannot create text")
	ErrCannotUpdateText          = errors.New("cannot update text")
	ErrCannotDeleteText          = errors.New("cannot update text")
	ErrCannotListText            = errors.New("cannot list texts")
	ErrCannotGetText             = errors.New("cannot get text details")
	ErrRequesterIsNotOwner       = errors.New("no permission, only text owner can do this")
	ErrCannotGetSummary          = errors.New("cannot get summary")
	ErrRequesterCannotModify     = errors.New("no permission to modify this text")
	ErrRequesterCannotRead       = errors.New("no permission to read this text, only owner or collaborator can read this text")
)
