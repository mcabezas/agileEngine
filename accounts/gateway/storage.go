package gateway

import (
	"context"
	"fmt"
	"sync"

	"github.com/mcabezas/agileEngine/accounts/models"
)

type Storage interface {
	Create(ctx context.Context, acc *models.Account) error
	Update(ctx context.Context, acc *models.Account) error
	Get(ctx context.Context, accountID string) (*models.Account, error)
}

type storage struct {
	accounts *sync.Map
}

func NewStorage() Storage {
	return &storage{accounts: &sync.Map{}}
}

func (s *storage) Create(_ context.Context, acc *models.Account) error {
	s.accounts.Store(acc.ID, acc)
	return nil
}

func (s *storage) Update(_ context.Context, acc *models.Account) error {
	s.accounts.Store(acc.ID, acc)
	return nil
}

func (s *storage) Get(_ context.Context, accountID string) (*models.Account, error) {
	acc, ok := s.accounts.Load(accountID)
	if !ok {
		return nil, fmt.Errorf("account not found")
	}
	account, ok := acc.(*models.Account)
	if !ok {
		return nil, fmt.Errorf("account not found")
	}
	return account, nil
}
