package entity

import "errors"

var (
	ErrContentNotValid     = errors.New("content cannot be blank")
	ErrUserIdNotValid      = errors.New("user id is not valid")
	ErrSummaryNotFound     = errors.New("summary not found")
	ErrCannotCreateSummary = errors.New("cannot create summary")
	ErrCannotUpdateSummary = errors.New("cannot update summary")
	ErrCannotDeleteSummary = errors.New("cannot delete summary")
	ErrCannotListSummary   = errors.New("cannot list summaries")
	ErrCannotGetSummary    = errors.New("cannot get summary details")
	ErrRequesterIsNotOwner = errors.New("no permission, only summary owner can do this")
)
