package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := biz.ValidateRegistrationInput(ctx, data); err != nil {
		return err
	}

	salt, hashedPassword, err := biz.ProcessPassword(data.Password)
	if err != nil {
		return err
	}

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, hashedPassword)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		biz.CompensateUserCreation(ctx, newUserId)
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	if err := biz.SendVerificationEmail(ctx, data.Email); err != nil {
		biz.CompensateAuthCreation(ctx, data.Email)
		biz.CompensateUserCreation(ctx, newUserId)
		return err
	}

	return nil
}
