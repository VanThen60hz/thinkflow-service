package business

import (
	"context"
	"fmt"

	"github.com/VanThen60hz/service-context/core"
	"github.com/golang-jwt/jwt/v5"
)

func (biz *business) IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	if accessToken == "" {
		return nil, core.ErrUnauthorized.WithError("access token is empty")
	}

	if biz.jwtProvider == nil {
		return nil, core.ErrInternalServerError.WithError("jwt provider is not initialized")
	}

	blacklistKey := fmt.Sprintf("blacklist:token:%s", accessToken)
	_, err := biz.redisClient.Get(ctx, blacklistKey)
	if err == nil {
		return nil, core.ErrUnauthorized.WithError("token has been revoked")
	}

	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)
	if err != nil {
		return nil, core.ErrUnauthorized.WithDebug(err.Error())
	}

	return claims, nil
}
