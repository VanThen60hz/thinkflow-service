package entity

import "strings"

type EmailVerificationRequest struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}

func (r *EmailVerificationRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if !emailIsValid(r.Email) {
		return ErrEmailIsNotValid
	}

	r.OTP = strings.TrimSpace(r.OTP)
	if len(r.OTP) != 6 {
		return ErrInvalidOTP
	}

	return nil
}
