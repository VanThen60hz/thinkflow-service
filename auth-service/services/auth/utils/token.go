package utils

import (
	"context"

	"thinkflow-service/common"
	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/google/uuid"
)

// GenerateToken creates a JWT token for the given user ID
func GenerateToken(ctx context.Context, jwtProvider common.JWTProvider, userId int) (*entity.TokenResponse, error) {
	uid := core.NewUID(uint32(userId), 1, 1)
	sub := uid.String()
	tid := uuid.New().String()

	tokenStr, expSecs, err := jwtProvider.IssueToken(ctx, tid, sub)
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

// JWTProvider interface for JWT token operations
type JWTProvider interface {
	IssueToken(ctx context.Context, tid, sub string) (string, int64, error)
}
