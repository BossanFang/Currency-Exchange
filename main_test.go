package main

import (
	"currency_exchange/middleware"
	"currency_exchange/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type ExchangeResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func TestExchangeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.ValidateInput())
	exchangeService := service.NewCurrencyExchangeService()
	router.GET("/exchange", exchangeService.ExchangeHandler)

	tests := []struct {
		query        string
		expectedCode int
		expectedBody ExchangeResponse
	}{
		// 無效的 source
		{"?source=XYZ&target=USD&amount=1000", http.StatusBadRequest, ExchangeResponse{}},
		// 無效的 target
		{"?source=USD&target=XYZ&amount=1000", http.StatusBadRequest, ExchangeResponse{}},
		// 輸入金額為非數字金額
		{"?source=USD&target=JPY&amount=abc", http.StatusBadRequest, ExchangeResponse{}},
		// 數字金額為負數
		{"?source=USD&target=JPY&amount=-1000", http.StatusBadRequest, ExchangeResponse{}},
		// 有效的 source 和 target，輸入金額為0
		{"?source=USD&target=JPY&amount=0", http.StatusOK, ExchangeResponse{"success", "0.00"}},
		// 有效的 source 和 target，輸入金額有千分位
		{"?source=USD&target=JPY&amount=1,525", http.StatusOK, ExchangeResponse{"success", "170,496.53"}},
		// 有效的 source 和 target，輸入金額為整數
		{"?source=USD&target=JPY&amount=1000", http.StatusOK, ExchangeResponse{"success", "111,801.00"}},
		// 有效的 source 和 target，輸入金額有小數
		{"?source=USD&target=JPY&amount=1000.456123", http.StatusOK, ExchangeResponse{"success", "111,852.43"}},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/exchange"+test.query, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedCode)
		}

		if test.expectedCode == http.StatusOK {
			var response ExchangeResponse
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Could not unmarshal response: %v", err)
			}

			if response != test.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", response, test.expectedBody)
			}
		}
	}
}
