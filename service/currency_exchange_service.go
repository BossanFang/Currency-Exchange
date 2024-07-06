package service

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ExchangeRates struct {
	Table map[string]map[string]float64 `json:"table"`
}

type CurrencyExchangeService struct {
	Rates ExchangeRates
}

func NewCurrencyExchangeService() *CurrencyExchangeService {
	return &CurrencyExchangeService{
		Rates: ExchangeRates{
			Table: map[string]map[string]float64{
				"TWD": {"TWD": 1, "JPY": 3.669, "USD": 0.03281},
				"JPY": {"TWD": 0.26956, "JPY": 1, "USD": 0.00885},
				"USD": {"TWD": 30.444, "JPY": 111.801, "USD": 1},
			},
		},
	}
}

func (s *CurrencyExchangeService) formatAmount(amount float64) string {
	// 四捨五入到小數點第二位
	roundedAmount := math.Round(amount*100) / 100
	parts := strings.Split(strconv.FormatFloat(roundedAmount, 'f', 2, 64), ".")
	intPart := parts[0]
	decPart := parts[1]

	var result strings.Builder
	for i, digit := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(digit)
	}
	return result.String() + "." + decPart
}

func (s *CurrencyExchangeService) ExchangeHandler(c *gin.Context) {
	source := c.GetString("source")
	target := c.GetString("target")
	amount := c.GetFloat64("amount")

	sourceRates, sourceExists := s.Rates.Table[source]
	if !sourceExists {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Unsupported source currency"})
		return
	}

	rate, targetExists := sourceRates[target]
	if !targetExists {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Unsupported target currency"})
		return
	}

	convertedAmount := amount * rate
	formattedAmount := s.formatAmount(convertedAmount)

	c.JSON(http.StatusOK, gin.H{"msg": "success", "amount": formattedAmount})
}
