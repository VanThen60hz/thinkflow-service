package entity

import "strings"

type ResetPasswordRequest struct {
	Email       string `json:"email" form:"email"`
	OTP         string `json:"otp" form:"otp"`
	NewPassword string `json:"new_password" form:"new_password"`
}

func (r *ResetPasswordRequest) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if !emailIsValid(r.Email) {
		return ErrEmailIsNotValid
	}

	r.OTP = strings.TrimSpace(r.OTP)
	if len(r.OTP) != 6 {
		return ErrInvalidOTP
	}

	r.NewPassword = strings.TrimSpace(r.NewPassword)
	if err := checkPassword(r.NewPassword); err != nil {
		return err
	}

	return nil
}
