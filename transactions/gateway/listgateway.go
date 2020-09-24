package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type ListGateway interface {
	List(ctx context.Context, accountID string) ([]*models.Transaction, error)
}

func NewListGateway(l logs.Logger, s Storage) ListGateway {
	return &ListGtw{
		log:     l,
		Storage: s,
	}
}


