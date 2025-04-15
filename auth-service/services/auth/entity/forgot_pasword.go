package entity

import "strings"

type ForgotPasswordRequest struct {
	Email string `json:"email" form:"email"`
}

func (f *ForgotPasswordRequest) Validate() error {
	f.Email = strings.TrimSpace(f.Email)
	if !emailIsValid(f.Email) {
		return ErrEmailIsNotValid
	}
	return nil
}
