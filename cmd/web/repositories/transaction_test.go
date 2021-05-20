package repositories

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionRepository_GetWallet(t *testing.T) {
	setTestEnvironment()

	repository := NewTransactionRepository()

	wallet, _ := repository.GetWallet(1)

	assert.True(t, wallet.Balance.Equal(decimal.NewFromInt(20)))
}

func TestTransactionRepository_UpdateWallet(t *testing.T) {
	setTestEnvironment()

	repository := NewTransactionRepository()

	wallet, _ := repository.GetWallet(2)
	wallet.Balance = decimal.NewFromInt(30)
	repository.UpdateWallet(wallet)

	// wait for go routing to delete cache
	time.Sleep(1 * time.Second)

	wallet, _ = repository.GetWallet(2)
	assert.True(t, wallet.Balance.Equal(decimal.NewFromInt(30)))
}

func setTestEnvironment() {
	viper.Set("env", "test")
	viper.Set("database.host", "localhost:3305")
	viper.Set("database.pass", "root")
	viper.Set("database.user", "root")
	viper.Set("database.name", "challenge")

	viper.Set("cache.host","localhost")
	viper.Set("cache.port", "6378")
}
