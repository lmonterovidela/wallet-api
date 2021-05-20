package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wallet-api/cmd/web/models"
	"github.com/wallet-api/cmd/web/services"
	"github.com/wallet-api/exceptions"
	"net/http"
	"strconv"
)

type ITransactionHandler interface {
	GetBalance(c *gin.Context)
	Debit(c *gin.Context)
	Credit(c *gin.Context)
}

type TransactionHandler struct {
	transactionService services.ITransactionService
}

const (
	ErrorCodeInvalidParams         string = "invalid params"
	ErrorCodeInvalidParamsPositive string = "the amount must be positive"
)

func (handler *TransactionHandler) GetBalance(c *gin.Context) {
	walletIdParam := c.Params.ByName("wallet_id")
	walletId, err := strconv.Atoi(walletIdParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": ErrorCodeInvalidParams})
		return
	}
	balance, err := handler.transactionService.GetBalance(walletId)
	if err != nil {
		handlerException(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (handler *TransactionHandler) Debit(c *gin.Context) {
	walletIdParam := c.Params.ByName("wallet_id")
	walletId, err := strconv.Atoi(walletIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ErrorCodeInvalidParams})
		return
	}

	var walletRequest models.WalletRequest
	if err = c.Bind(&walletRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ErrorCodeInvalidParams})
		return
	}

	if err := handler.transactionService.Debit(walletId, walletRequest.Amount); err != nil {
		handlerException(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *TransactionHandler) Credit(c *gin.Context) {
	walletIdParam := c.Params.ByName("wallet_id")
	walletId, err := strconv.Atoi(walletIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ErrorCodeInvalidParams})
		return
	}

	var walletRequest models.WalletRequest
	if err = c.Bind(&walletRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": ErrorCodeInvalidParams})
		return
	}

	if err := handler.transactionService.Credit(walletId, walletRequest.Amount); err != nil {
		handlerException(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func handlerException(c *gin.Context, err error) {
	switch err.(type) {
	case *exceptions.NotFoundException:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	case *exceptions.ForbiddenException:
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	default:
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func NewTransactionHandler() ITransactionHandler {
	return &TransactionHandler{
		transactionService: services.NewTransactionService(),
	}
}
