package entity

import (
	"strings"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

type User struct {
	core.SQLModel             // in practice, we should not embed this struct
	FirstName     string      `json:"first_name" gorm:"column:first_name" db:"first_name"`
	LastName      string      `json:"last_name" gorm:"column:last_name" db:"last_name"`
	Email         string      `json:"email" gorm:"column:email" db:"email"`
	Phone         string      `json:"phone" gorm:"column:phone" db:"phone"`
	AvatarId      int         `json:"-" gorm:"column:avatar_id" db:"avatar_id"`
	Avatar        *core.Image `json:"avatar" gorm:"-" db:"-"`
	Gender        Gender      `json:"gender" gorm:"column:gender" db:"gender"`
	SystemRole    SystemRole  `json:"system_role" gorm:"column:system_role" db:"system_role"`
	Status        Status      `json:"status" gorm:"column:status" db:"status"`
}

func NewUser(firstName, lastName, email string) User {
	return User{
		SQLModel:   core.NewSQLModel(),
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Phone:      "",
		Avatar:     nil,
		Gender:     GenderUnknown,
		SystemRole: RoleUser,
		Status:     StatusPendingVerify,
	}
}

func (User) TableName() string { return "users" }

func (u *User) Validate() error {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)

	if err := checkFirstName(u.FirstName); err != nil {
		return err
	}

	if err := checkLastName(u.LastName); err != nil {
		return err
	}

	if !emailIsValid(u.Email) {
		return ErrEmailIsNotValid
	}

	return nil
}

func (u *User) Mask() {
	u.SQLModel.Mask(common.MaskTypeUser)

	if av := u.Avatar; av != nil {
		av.Mask(common.MaskTypeImage)
	}
}
