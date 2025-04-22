package business

import (
	"context"
	"fmt"
)

func (biz *business) CompensateAuthCreation(ctx context.Context, email string) {
	if err := biz.repository.DeleteAuth(ctx, email); err != nil {
		fmt.Printf("Failed to compensate auth creation: %v\n", err)
	}
}
