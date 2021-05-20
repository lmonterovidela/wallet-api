package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet-api/cmd/web/handlers"
	"net/http"
)

func Routes(r *gin.Engine) {

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	transactionHandler := handlers.NewTransactionHandler()
	r.GET("/api/v1/wallets/:wallet_id/balance", transactionHandler.GetBalance)
	r.POST("/api/v1/wallets/:wallet_id/debit", transactionHandler.Debit)
	r.POST("/api/v1/wallets/:wallet_id/credit", transactionHandler.Credit)
}
