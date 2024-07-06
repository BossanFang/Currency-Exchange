package middleware

import "github.com/gin-gonic/gin"

func ValidateInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate source, target, amount
	}
}
