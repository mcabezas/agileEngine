package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/models"
	"github.com/mcabezas/agileEngine/internal/logs"
)

type GetGateway interface {
	Get(ctx context.Context, accountID string) (*models.Account, error)
}

func NewGetGateway(l logs.Logger, s Storage) GetGateway {
	return &GetGtw{
		log:     l,
		Storage: s,
	}
}
