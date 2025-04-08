package common

import (
	"github.com/gin-gonic/gin"
)

func IsHTTPS(c *gin.Context) bool {
    return c.Request.TLS != nil || c.Request.Header.Get("X-Forwarded-Proto") == "https"
}
