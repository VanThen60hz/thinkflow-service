package middleware

import (
	"context"
	"fmt"
	"log"

	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
	"github.com/gin-gonic/gin"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (sub string, tid string, err error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(c *gin.Context) {
		log.Printf("Processing request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Request headers: %v", c.Request.Header)
		log.Printf("Cookies: %v", c.Request.Cookies())

		var token string
		var err error

		// Try to get token from cookie first
		token, err = c.Cookie("accessToken")
		if err != nil {
			// If not in cookie, try Authorization header
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				log.Printf("No access token found in cookie or Authorization header")
				core.WriteErrorResponse(c, core.ErrUnauthorized.WithError("missing access token"))
				c.Abort()
				return
			}

			// Extract token from Authorization header
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				token = authHeader[7:]
			} else {
				log.Printf("Invalid Authorization header format")
				core.WriteErrorResponse(c, core.ErrUnauthorized.WithError("invalid authorization header format"))
				c.Abort()
				return
			}
			log.Printf("Found access token in Authorization header")
		} else {
			log.Printf("Found access token in cookie")
		}

		sub, tid, err := ac.IntrospectToken(c.Request.Context(), token)
		if err != nil {
			log.Printf("Token validation failed: %v", err)
			core.WriteErrorResponse(c, core.ErrUnauthorized.WithDebug(err.Error()))
			c.Abort()
			return
		}

		log.Printf("Token validated successfully for user: %s, tenant: %s", sub, tid)

		requester := core.NewRequester(sub, tid)
		c.Set(common.RequesterKey, requester)

		ctx := context.WithValue(c.Request.Context(), common.RequesterKey, requester)
		c.Request = c.Request.WithContext(ctx)

		fmt.Println("Request authenticated", "sub", sub, "tid", tid)
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
