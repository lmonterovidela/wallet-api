package services

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wallet-api/cmd/web/models"
	mocks "github.com/wallet-api/mocks/repositories"
	"testing"
)

func TestTransactionService_GetBalance(t *testing.T) {

	repositoryMock := &mocks.RepositoryMock{}

	type args struct {
		walletId int
	}

	tests := []struct {
		name        string
		initMocks   func()
		args        args
		assertMocks func(*testing.T)
		assertError func(*testing.T, error)
		assertFunc  func(*testing.T, decimal.Decimal)
	}{
		{
			name: "Success - repository response ok",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{
						Balance: decimal.NewFromInt(222),
					}, nil).Once()
			},
			args: args{
				walletId: 1,
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
			assertFunc: func(t *testing.T, balance decimal.Decimal) {
				assert.Equal(t, decimal.NewFromInt(222), balance)
			},
		},
		{
			name: "Error - repository response err",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{}, errors.New("some error")).Once()
			},
			args: args{
				walletId: 1,
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
			assertFunc: func(t *testing.T, balance decimal.Decimal) {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			service := TransactionService{
				transactionRepository: repositoryMock,
			}

			balance, err := service.GetBalance(tt.args.walletId)
			tt.assertMocks(t)
			tt.assertError(t, err)
			tt.assertFunc(t, balance)
		})
	}
}

func TestTransactionService_Debit(t *testing.T) {

	repositoryMock := &mocks.RepositoryMock{}

	type args struct {
		walletId int
		amount   decimal.Decimal
	}

	tests := []struct {
		name        string
		initMocks   func()
		args        args
		assertMocks func(*testing.T)
		assertError func(*testing.T, error)
	}{
		{
			name: "Success - debit ok",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{
						Balance: decimal.NewFromInt(200),
					}, nil).Once()
				repositoryMock.On("UpdateWallet", mock.Anything).
					Return(nil).Once()
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(12),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
		},
		{
			name: "Error - repository response err",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{}, errors.New("some error")).Once()
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(12),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
		{
			name: "Error - negative balance",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{Balance: decimal.NewFromInt(200)}, nil).Once()
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(240),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
		{
			name: "Error - negative amount",
			initMocks: func() {
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(-1),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			service := TransactionService{
				transactionRepository: repositoryMock,
			}

			err := service.Debit(tt.args.walletId, tt.args.amount)
			tt.assertMocks(t)
			tt.assertError(t, err)
		})
	}
}

func TestTransactionService_Credit(t *testing.T) {

	repositoryMock := &mocks.RepositoryMock{}

	type args struct {
		walletId int
		amount   decimal.Decimal
	}

	tests := []struct {
		name        string
		initMocks   func()
		args        args
		assertMocks func(*testing.T)
		assertError func(*testing.T, error)
	}{
		{
			name: "Success - credit ok",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{
						Balance: decimal.NewFromInt(200),
					}, nil).Once()
				repositoryMock.On("UpdateWallet", mock.Anything).
					Return(nil).Once()
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(12),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.Nil(t, e)
			},
		},
		{
			name: "Error - repository response err",
			initMocks: func() {
				repositoryMock.On("GetWallet", 1).
					Return(models.Wallet{}, errors.New("some error")).Once()
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(12),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
		{
			name: "Error - negative amount",
			initMocks: func() {
			},
			args: args{
				walletId: 1,
				amount:   decimal.NewFromInt(-1),
			},
			assertMocks: func(t *testing.T) {
				repositoryMock.AssertExpectations(t)
			},
			assertError: func(t *testing.T, e error) {
				assert.NotNil(t, e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initMocks()
			service := TransactionService{
				transactionRepository: repositoryMock,
			}

			err := service.Credit(tt.args.walletId, tt.args.amount)
			tt.assertMocks(t)
			tt.assertError(t, err)
		})
	}
}
