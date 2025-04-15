package middleware

import (
	"context"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("accessToken")
		if err != nil {
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithError("missing access token in cookie"))
			c.Abort()
			return
		}

		sub, tid, err := ac.IntrospectToken(c.Request.Context(), token)
		if err != nil {
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithDebug(err.Error()))
			c.Abort()
			return
		}

		c.Set(string(core.KeyRequester), core.NewRequester(sub, tid))

		c.Next()
	}
}

// func extractTokenFromHeaderString(s string) (string, error) {
// 	parts := strings.Split(s, " ")
// 	//"Authorization" : "Bearer {token}"

// 	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" || strings.TrimSpace(parts[1]) == "null" {
// 		return "", core.ErrUnauthorized.WithError("missing access token")
// 	}

// 	return parts[1], nil
// }
