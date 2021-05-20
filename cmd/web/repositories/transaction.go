package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/wallet-api/cmd/web/models"
	"github.com/wallet-api/exceptions"
	"github.com/wallet-api/infrastructure"
	"strconv"
)

type ITransactionRepository interface {
	GetWallet(walletId int) (models.Wallet, error)
	UpdateWallet(wallet models.Wallet) error
}

type TransactionRepository struct {
	dbProvider    *gorm.DB
	cacheProvider infrastructure.ICacheProvider
}

const walletKey string = "wallet_%d"
const walletNotFound string = "wallet with id=%d not found"

func (repository *TransactionRepository) GetWallet(walletId int) (models.Wallet, error) {
	// find in cache
	wallet, err := repository.getWalletFromCache(walletId)
	if err == nil {
		return wallet, nil
	}

	//find in database
	status := repository.dbProvider.First(&wallet, walletId)

	if status.Error != nil {
		return wallet, status.Error
	}

	if wallet.ID == 0 {
		return wallet, exceptions.NewNotFoundException(walletNotFound, strconv.Itoa(walletId))
	}

	// save in cache
	j, err := json.Marshal(wallet)
	if err != nil {
		return wallet, nil
	}
	go repository.cacheProvider.Set(fmt.Sprintf(walletKey, walletId), j, 0)

	return wallet, nil
}

func (repository *TransactionRepository) UpdateWallet(wallet models.Wallet) error {
	status := repository.dbProvider.Save(&wallet)

	go repository.cacheProvider.Set(fmt.Sprintf(walletKey, wallet.ID), nil, 0)

	return status.Error
}

func (repository *TransactionRepository) getWalletFromCache(walletId int) (models.Wallet, error) {
	// find on cache
	result, err := repository.cacheProvider.Get(fmt.Sprintf(walletKey, walletId))
	if err != nil {
		return models.Wallet{}, err
	}
	// if not found key
	if result == "" {
		return models.Wallet{}, errors.New("not found")
	}

	var wallet *models.Wallet
	err = json.Unmarshal([]byte(result), &wallet)
	if err != nil {
		logrus.Errorf("couldn't unmarshal wallet from cache: %v", err)
		return models.Wallet{}, fmt.Errorf("couldn't unmarshal wallet from cache: %v", err)
	}

	return *wallet, nil
}

func NewTransactionRepository() ITransactionRepository {
	return &TransactionRepository{
		dbProvider:    infrastructure.ConnectDatabase(),
		cacheProvider: infrastructure.NewCacheClient(),
	}
}
