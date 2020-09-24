package gateway

import (
	"context"
	"sync"
	"time"

	"github.com/mcabezas/agileEngine/internal"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type Storage interface {
	List(ctx context.Context, accountID string) ([]*models.Transaction, error)
	Get(ctx context.Context, transactionID string) (*models.Transaction, error)
	Create(_ context.Context, cmd *models.CreateTransactionCMD) (*models.Transaction, error)
}

type storage struct {
	transactions *sync.Map
}

func NewStorage() Storage {
	return &storage{transactions: &sync.Map{}}
}

func (s *storage) List(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	var trxs []*models.Transaction
	s.transactions.Range(func(k, v interface{}) bool {
		if transaction, ok := v.(*models.Transaction); ok {
			if transaction.AccountID == accountID {
				trxs = append(trxs, transaction)
			}
		}
		return true
	})
	return trxs, nil
}

func (s *storage) Get(ctx context.Context, transactionID string) (*models.Transaction, error) {
	trx, _ := s.transactions.Load(transactionID)
	return trx.(*models.Transaction), nil
}

func (s *storage) Create(_ context.Context, cmd *models.CreateTransactionCMD) (*models.Transaction, error) {
	transaction := &models.Transaction{
		ID:              internal.UUID(),
		AccountID:       cmd.AccountID,
		TransactionType: cmd.TransactionType,
		Amount:          cmd.Amount,
		CreatedAt:       time.Now().Unix(),
	}
	s.transactions.Store(transaction.ID, transaction)
	return transaction, nil
}
