package main

import (
	"currency_exchange/middleware"
	"currency_exchange/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.ValidateInput())
	exchangeService := service.NewCurrencyExchangeService()
	r.GET("/exchange", exchangeService.ExchangeHandler)
	r.Run(":8080")
}
