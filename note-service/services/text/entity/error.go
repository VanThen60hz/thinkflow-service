package entity

import "errors"

var (
	ErrTextContentCannotNotBlank = errors.New("text content cannot be blank")
	ErrUserIdNotValid            = errors.New("user id is not valid")
	ErrTextNotFound              = errors.New("text not found")
	ErrCannotCreateText          = errors.New("cannot create text")
	ErrCannotUpdateText          = errors.New("cannot update text")
	ErrCannotDeleteText          = errors.New("cannot update text")
	ErrCannotListText            = errors.New("cannot list texts")
	ErrCannotGetText             = errors.New("cannot get text details")
	ErrRequesterIsNotOwner       = errors.New("no permission, only text owner can do this")
)
