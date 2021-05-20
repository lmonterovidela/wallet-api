package services

import (
	"github.com/shopspring/decimal"
	"github.com/wallet-api/cmd/web/repositories"
	"github.com/wallet-api/exceptions"
)

type ITransactionService interface {
	GetBalance(walletId int) (decimal.Decimal, error)
	Debit(walletId int, amount decimal.Decimal) error
	Credit(walletId int, amount decimal.Decimal) error
}

type TransactionService struct {
	transactionRepository repositories.ITransactionRepository
}

const ErrorCodeInvalidParamsPositive string = "the amount must be positive"
const ErrorCodeInvalid string = "a wallet balance cannot go below 0."

func (service *TransactionService) GetBalance(walletId int) (decimal.Decimal, error) {
	wallet, err := service.transactionRepository.GetWallet(walletId)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return wallet.Balance, err
}

func (service *TransactionService) Debit(walletId int, amount decimal.Decimal) error {
	//If it were a real feature, it would be necessary to have a lock or do the two queries in formal transaction
	if !amount.IsPositive() {
		return exceptions.NewForbiddenException("operation not allowed: %s", ErrorCodeInvalidParamsPositive)
	}

	wallet, err := service.transactionRepository.GetWallet(walletId)
	if err != nil {
		return err
	}
	if wallet.Balance.LessThanOrEqual(amount) {
		return exceptions.NewForbiddenException("operation not allowed: %s", ErrorCodeInvalid)
	}

	wallet.Balance = wallet.Balance.Sub(amount)
	return service.transactionRepository.UpdateWallet(wallet)
}

func (service *TransactionService) Credit(walletId int, amount decimal.Decimal) error {
	//If it were a real feature, it would be necessary to have a lock or do the two queries in formal transaction
	if !amount.IsPositive() {
		return exceptions.NewForbiddenException("operation not allowed: %s", ErrorCodeInvalidParamsPositive)
	}

	wallet, err := service.transactionRepository.GetWallet(walletId)
	if err != nil {
		return err
	}

	wallet.Balance = wallet.Balance.Add(amount)
	return service.transactionRepository.UpdateWallet(wallet)
}

func NewTransactionService() ITransactionService {
	return &TransactionService{
		transactionRepository: repositories.NewTransactionRepository(),
	}
}
