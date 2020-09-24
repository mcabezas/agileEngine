package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type ListGtw struct {
	log logs.Logger
	Storage
}

func (g *ListGtw) List(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	return g.Storage.List(ctx, accountID)
}