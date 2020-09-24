package models

import (
	"sync"

	"github.com/mcabezas/agileEngine/internal/money"
)

type Account struct {
	ID      string
	balance money.Money
	rwMutex *sync.RWMutex
}

func NewAccount(id string, balance money.Money) *Account {
	return &Account{
		ID:      id,
		balance: balance,
		rwMutex: &sync.RWMutex{},
	}
}

func (acc *Account) Balance() money.Money {
	acc.rwMutex.RLock()
	defer acc.rwMutex.RUnlock()
	return acc.balance
}

func (acc *Account) LockBalance() money.Money{
	acc.rwMutex.Lock()
	return acc.balance
}

func (acc *Account) UnlockBalance() {
	acc.rwMutex.Unlock()
}