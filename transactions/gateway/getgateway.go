package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type GetGateway interface {
	Get(ctx context.Context, transactionID string) (*models.Transaction, error)
}

func NewGetGateway(l logs.Logger, s Storage) GetGateway {
	return &GetGtw{
		log:     l,
		Storage: s,
	}
}


