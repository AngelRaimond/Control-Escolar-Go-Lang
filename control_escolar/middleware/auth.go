package middleware

import "github.com/gin-gonic/gin"

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("x-api-key")
		if apikey == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "API key is required",
			})
			return
		}
		if apikey != "jkhdkasjhkdhkajhdjkhilh387uuij3ljk" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid API key",
			})
			return
		}
		c.Next()
	}
}
