package gateway

import (
	"context"
	"testing"

	accGateway "github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
	"github.com/mcabezas/agileEngine/transactions/models"
	"github.com/stretchr/testify/assert"
)

func TestNewCanNotCreateNegativeBalances(t *testing.T) {
	accounts, accID := createAccount()
	storage := newStorage()
	createG := NewCreateGateway(logs.NewSugaredLogger(), storage, accounts)
//	listG := NewListGateway(logs.NewSugaredLogger(), storage)
	_, err := createG.Create(context.Background(), models.CreateTransactionCMD{
		AccountID:       accID,
		TransactionType: models.Debit,
		Amount:          money.Money{
			Amount: money.Amount(10),
			Currency: money.DefaultCurrency(),
		},
	})
	assert.NotNil(t, err)
}

func TestCanListTransactions(t *testing.T) {
	accounts, accID := createAccount()
	storage := newStorage()
	createG := NewCreateGateway(logs.NewSugaredLogger(), storage, accounts)
	listG := NewListGateway(logs.NewSugaredLogger(), storage)
	_, _ = createG.Create(context.Background(), models.CreateTransactionCMD{
		AccountID:       accID,
		TransactionType: models.Credit,
		Amount:          money.Money{
			Amount: money.Amount(10),
			Currency: money.DefaultCurrency(),
		},
	})
	_, _ = createG.Create(context.Background(), models.CreateTransactionCMD{
		AccountID:       accID,
		TransactionType: models.Credit,
		Amount:          money.Money{
			Amount: money.Amount(20),
			Currency: money.DefaultCurrency(),
		},
	})
	_, _ = createG.Create(context.Background(), models.CreateTransactionCMD{
		AccountID:       accID,
		TransactionType: models.Debit,
		Amount:          money.Money{
			Amount: money.Amount(5),
			Currency: money.DefaultCurrency(),
		},
	})
	list, _ := listG.List(context.Background(), accID)
	assert.True(t, len(list) == 3)
}

func createAccount() (*accGateway.AccountGateway, string) {
	storage := accGateway.NewStorage()
	createG := accGateway.NewCreateGateway(logs.NewSugaredLogger(), storage)
	createdID, _ := createG.Create(context.Background())
	return &accGateway.AccountGateway{
		GetGateway:    accGateway.NewGetGateway(logs.NewSugaredLogger(), storage),
		CreateGateway: createG,
		UpdateGateway: accGateway.NewUpdateGateway(logs.NewSugaredLogger(), storage),
	}, createdID
}