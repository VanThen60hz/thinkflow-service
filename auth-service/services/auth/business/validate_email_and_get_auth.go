package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) ValidateEmailAndGetAuth(ctx context.Context, email string) (*entity.Auth, error) {
	auth, err := biz.repository.GetAuth(ctx, email)
	if err != nil {
		if err == core.ErrRecordNotFound {
			return nil, core.ErrBadRequest.WithError(entity.ErrEmailNotFound.Error())
		}
		return nil, core.ErrInternalServerError.WithDebug(err.Error())
	}
	return auth, nil
}
