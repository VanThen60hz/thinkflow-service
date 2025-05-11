package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	allowedOrigins := []string{
		"http://localhost:3001",
		"http://localhost:3002",
		"http://42.113.255.139:5500",
		"http://127.0.0.1:5500",
		"http://118.70.192.62:3001",
		"http://118.70.192.62:3002",
		"https://thinkflow-web.vercel.app",
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Printf("Received request from origin: %s\n", origin)
		fmt.Printf("Request method: %s\n", c.Request.Method)
		fmt.Printf("Request headers: %v\n", c.Request.Header)
		
		if isAllowedOrigin(origin, allowedOrigins) {
			fmt.Printf("Origin %s is allowed\n", origin)
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "ETag, Accept-Ranges, Content-Encoding, Content-Length, Content-Range")
		} else {
			fmt.Printf("Origin %s is NOT allowed\n", origin)
		}

		if c.Request.Method == "OPTIONS" {
			fmt.Println("Handling OPTIONS request")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func isAllowedOrigin(origin string, allowed []string) bool {
	for _, o := range allowed {
		if strings.EqualFold(o, origin) {
			return true
		}
	}
	return false
}
