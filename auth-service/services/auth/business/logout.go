package business

import (
	"context"
	"fmt"
	"time"

	"github.com/VanThen60hz/service-context/core"
)

func (biz *business) Logout(ctx context.Context, accessToken string) error {
	claims, err := biz.jwtProvider.ParseToken(ctx, accessToken)
	if err != nil {
		return core.ErrUnauthorized.WithDebug(err.Error())
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	ttl := time.Until(expTime.Time)
	if ttl <= 0 {
		return nil
	}

	blacklistKey := fmt.Sprintf("blacklist:token:%s", accessToken)
	err = biz.redisClient.Set(ctx, blacklistKey, "revoked", ttl)
	if err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	return nil
}
