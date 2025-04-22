package business

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

func (biz *business) GenerateToken(ctx context.Context, userId int) (*entity.TokenResponse, error) {
	uid := core.NewUID(uint32(userId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := biz.jwtProvider.IssueToken(ctx, tid, sub)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(entity.ErrLoginFailed.Error()).WithDebug(err.Error())
	}

	return &entity.TokenResponse{
		AccessToken: entity.Token{
			Token:     tokenStr,
			ExpiredIn: int(expSecs),
		},
	}, nil
}
