package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tabed23/user-boilerplate-auth/utils"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
