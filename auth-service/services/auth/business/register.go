package business

import (
	"context"

	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	if err := utils.ValidateRegistrationInput(ctx, biz.repository, data); err != nil {
		return err
	}

	salt, hashedPassword, err := utils.ProcessPassword(biz.hasher, data.Password)
	if err != nil {
		return err
	}

	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, hashedPassword)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		utils.CompensateUserCreation(ctx, biz.userRepository, newUserId)
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	if err := utils.SendVerificationEmail(ctx, biz.redisClient, biz.emailService, data.Email); err != nil {
		utils.CompensateAuthCreation(ctx, biz.repository, data.Email)
		utils.CompensateUserCreation(ctx, biz.userRepository, newUserId)
		return err
	}

	return nil
}
