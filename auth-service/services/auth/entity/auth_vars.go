package entity

import "strings"

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (ad *AuthEmailPassword) Validate() error {
	ad.Email = strings.TrimSpace(ad.Email)

	if !emailIsValid(ad.Email) {
		return ErrEmailIsNotValid
	}

	ad.Password = strings.TrimSpace(ad.Password)

	if err := checkPassword(ad.Password); err != nil {
		return err
	}

	return nil
}

type AuthRegister struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	AuthEmailPassword
}

func (ar *AuthRegister) Validate() error {
	if err := ar.AuthEmailPassword.Validate(); err != nil {
		return err
	}

	ar.FirstName = strings.TrimSpace(ar.FirstName)

	if err := checkFirstName(ar.FirstName); err != nil {
		return err
	}

	ar.LastName = strings.TrimSpace(ar.LastName)

	if err := checkLastName(ar.LastName); err != nil {
		return err
	}

	return nil
}

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int `json:"expire_in"`
}

type TokenResponse struct {
	AccessToken Token `json:"access_token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refresh_token,omitempty"`
}

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
