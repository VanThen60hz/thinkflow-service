package entity

import "strings"

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
