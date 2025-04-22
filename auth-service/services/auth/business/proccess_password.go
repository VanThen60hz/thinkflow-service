package business

import (
	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ProcessPassword(password string) (salt, hashedPassword string, err error) {
	salt, err = biz.hasher.RandomStr(16)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	hashedPassword, err = biz.hasher.HashPassword(salt, password)
	if err != nil {
		return "", "", core.ErrInternalServerError.WithDebug(err.Error())
	}

	return salt, hashedPassword, nil
}
