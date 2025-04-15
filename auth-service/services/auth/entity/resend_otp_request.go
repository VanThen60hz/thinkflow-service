package entity

import "strings"

type ResendOTPRequest struct {
	Email string `json:"email" form:"email"`
}

func (r *ResendOTPRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if !emailIsValid(r.Email) {
		return ErrEmailIsNotValid
	}
	return nil
}
