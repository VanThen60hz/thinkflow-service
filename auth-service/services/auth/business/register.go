package business

import (
	"context"
	"fmt"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"
	"thinkflow-service/services/auth/utils"

	"github.com/VanThen60hz/service-context/core"
)

// Register implements the Saga Pattern for user registration
func (biz *business) Register(ctx context.Context, data *entity.AuthRegister) error {
	// Step 1: Validate input data
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	// Step 2: Check if email already exists
	_, err := biz.repository.GetAuth(ctx, data.Email)
	if err == nil {
		return core.ErrBadRequest.WithError(entity.ErrEmailHasExisted.Error())
	} else if err != core.ErrRecordNotFound {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Step 3: Generate salt and hash password
	salt, err := biz.hasher.RandomStr(16)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	passHashed, err := biz.hasher.HashPassword(salt, data.Password)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Step 4: Create user in user service
	newUserId, err := biz.userRepository.CreateUser(ctx, data.FirstName, data.LastName, data.Email)
	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Compensation function to delete user if auth creation fails
	compensateUserCreation := func() {
		// We don't return error here as this is a compensation step
		// Log the error but don't propagate it
		if err := biz.userRepository.DeleteUser(ctx, newUserId); err != nil {
			// For now, just log the error since DeleteUser is not fully implemented
			fmt.Printf("Failed to compensate user creation: %v\n", err)
			// TODO: Implement proper user deletion in user service
		}
	}

	// Step 5: Create auth record
	newAuth := entity.NewAuthWithEmailPassword(newUserId, data.Email, salt, passHashed)
	if err := biz.repository.AddNewAuth(ctx, &newAuth); err != nil {
		// Compensation: Delete the user we just created
		compensateUserCreation()
		return core.ErrInternalServerError.WithError(entity.ErrCannotRegister.Error()).WithDebug(err.Error())
	}

	// Step 6: Generate and send OTP
	otp := core.GenerateOTP()
	err = utils.SendOTPEmail(
		ctx,
		biz.redisClient,
		biz.emailService,
		data.Email,
		otp,
		common.EmailVerifyOTPSubject,
		"Email Verification",
		"Thanks for signing up! Please use the OTP below to verify your email:",
		"email verification",
	)
	if err != nil {
		// Compensation: Delete both auth and user
		if err := biz.repository.DeleteAuth(ctx, data.Email); err != nil {
			fmt.Printf("Failed to compensate auth creation: %v\n", err)
		}
		compensateUserCreation()
		return err
	}

	return nil
}
