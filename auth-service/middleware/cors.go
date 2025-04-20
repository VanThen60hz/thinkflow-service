package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", 
		"http://localhost:3001, http://localhost:3002, http://118.70.192.62:3001, http://118.70.192.62:3002, https://thinkflow-web.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, X-Requested-With, Authorization, Range")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Accept-Ranges, Content-Range")		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
