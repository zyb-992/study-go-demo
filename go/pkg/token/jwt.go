package token

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWT Parse jwt
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if claim := ParseToken(token); claim == nil {
		}
	}
}

func GenerateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GenerateToken()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "error when generate token",
			})
			c.Abort()
			return
		}

		c.Header("authorization", token)
	}
}
