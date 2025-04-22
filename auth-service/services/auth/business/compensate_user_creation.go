package business

import (
	"context"
	"fmt"
)

func (biz *business) CompensateUserCreation(ctx context.Context, userId int) {
	if err := biz.userRepository.DeleteUser(ctx, userId); err != nil {
		fmt.Printf("Failed to compensate user creation: %v\n", err)
	}
}