package utils

import (
	"context"
	"fmt"
)

func CompensateUserCreation(ctx context.Context, userRepository UserRepository, userId int) {
	if err := userRepository.DeleteUser(ctx, userId); err != nil {
		fmt.Printf("Failed to compensate user creation: %v\n", err)
	}
}

func CompensateAuthCreation(ctx context.Context, repository AuthRepository, email string) {
	if err := repository.DeleteAuth(ctx, email); err != nil {
		fmt.Printf("Failed to compensate auth creation: %v\n", err)
	}
}
