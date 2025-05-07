package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) RegisterWithUserId(ctx context.Context, data *entity.AuthRegister, newUserId int) error {
	salt, hashedPassword, err := biz.ProcessPassword(data.Password)
	if err != nil {
		return err
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, hashedPassword)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	if err := biz.SendVerificationEmail(ctx, data.Email); err != nil {
		return err
	}

	return nil
}
