package models

import (
	"github.com/mcabezas/agileEngine/internal/money"
)

type TransactionType string

type Transaction struct {
	ID              string
	AccountID       string
	TransactionType TransactionType
	Amount          money.Money
	CreatedAt       int64
}

type CreateTransactionCMD struct {
	AccountID       string
	TransactionType TransactionType
	Amount          money.Money
}

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)
