package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wallet-api/cmd/web/models"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetWallet(walletId int) (models.Wallet, error) {
	args := m.Called(walletId)
	err := args.Error(1)
	if args.Get(0) == nil {
		return models.Wallet{}, err
	}
	return args.Get(0).(models.Wallet), err
}

func (m *RepositoryMock) UpdateWallet(wallet models.Wallet) error {
	args := m.Called(wallet)
	return args.Error(0)
}
