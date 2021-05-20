package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type WalletRequest struct {
	Amount decimal.Decimal `json:"amount" binding:"required"`
}

type Wallet struct {
	gorm.Model
	Balance decimal.Decimal `json:"balance" sql:"type:decimal(20,8)"`
}
