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
