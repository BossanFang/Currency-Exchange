package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyExchangeService struct{}

func NewCurrencyExchangeService() *CurrencyExchangeService {
	return &CurrencyExchangeService{}
}

func (s *CurrencyExchangeService) ExchangeHandler(c *gin.Context) {
	// implementation
	c.JSON(http.StatusOK, gin.H{"msg": "success", "amount": 0})
}
