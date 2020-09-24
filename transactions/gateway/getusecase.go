package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type GetGtw struct {
	log      logs.Logger
	Storage
}

func (g *GetGtw) Get(ctx context.Context, transactionID string) (*models.Transaction, error) {
	return g.Storage.Get(ctx, transactionID)
}
