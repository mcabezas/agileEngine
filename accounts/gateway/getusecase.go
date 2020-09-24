package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/models"
	"github.com/mcabezas/agileEngine/internal/logs"
)

type GetGtw struct {
	log logs.Logger
	Storage
}

func (g *CreateGtw) Get(ctx context.Context, accountID string) (*models.Account, error) {
	return g.Storage.Get(ctx, accountID)
}
