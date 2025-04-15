package entity

import "errors"

var (
	ErrPasswordIsNotValid  = errors.New("password must have from 8 to 30 characters")
	ErrEmailIsNotValid     = errors.New("email is not valid")
	ErrEmailHasExisted     = errors.New("email has existed")
	ErrLoginFailed         = errors.New("email and password are not valid")
	ErrFirstNameIsEmpty    = errors.New("first name can not be blank")
	ErrFirstNameTooLong    = errors.New("first name too long, max character is 30")
	ErrLastNameIsEmpty     = errors.New("last name can not be blank")
	ErrLastNameTooLong     = errors.New("last name too long, max character is 30")
	ErrCannotRegister      = errors.New("cannot register")
	ErrInvalidOTP          = errors.New("invalid OTP format")
	ErrEmailNotFound       = errors.New("email not found")
	ErrInvalidOrExpiredOTP = errors.New("invalid or expired OTP")
	ErrEmailNotVerified    = errors.New("email address has not been verified. Please check your email for the verification code")
	ErrUserAlreadyVerified = errors.New("user is already verified")
	ErrUserStatusNotMatch  = errors.New("user status is not as expected")
	ErrCannotGetUser       = errors.New("cannot get user")
	ErrCannotCreateUser    = errors.New("cannot create user")
	ErrCannotUpdateUser    = errors.New("cannot update user")
	ErrCannotDeleteUser    = errors.New("cannot delete user")
	ErrDuplicateEmail      = errors.New("email already exists")
)
