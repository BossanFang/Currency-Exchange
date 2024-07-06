package middleware

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		source := c.Query("source")
		target := c.Query("target")
		amountStr := c.Query("amount")

		if source == "" || target == "" || amountStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Missing parameters"})
			c.Abort()
			return
		}

		c.Set("source", source)
		c.Set("target", target)

		amountStr = strings.ReplaceAll(amountStr, ",", "")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Invalid amount"})
			c.Abort()
			return
		}

		// 四捨五入到小數點第二位
		roundedAmount := math.Round(amount*100) / 100

		c.Set("amount", roundedAmount)
		c.Next()
	}
}
