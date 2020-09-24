package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type CreateGateway interface {
	Create(ctx context.Context, cmd *models.CreateTransactionCMD) (*models.Transaction, error)
}

func NewCreateGateway(l logs.Logger, s Storage, accounts *gateway.AccountGateway) CreateGateway {
	return &CreateGtw{
		log:     l,
		Storage: s,
		accounts: accounts,
	}
}
