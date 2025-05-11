package entity

import "errors"

var (
	ErrNotificationMessageIsEmpty   = errors.New("notification message can not be blank")
	ErrNotificationSenderIsRequired = errors.New("notification sender is required")
	ErrCannotCreateNotification     = errors.New("cannot create notification")
	ErrCannotGetNotification        = errors.New("cannot get notification")
	ErrRequesterIsNotReceivedUser   = errors.New("requester is not the receiver of this notification")
	ErrCannotDeleteNotification     = errors.New("cannot delete notification")
	ErrCannotGetUser                = errors.New("cannot get user")
)
