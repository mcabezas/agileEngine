package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/internal/logs"
)

type CreateGateway interface {
	Create(ctx context.Context) (string, error)
}

func NewCreateGateway(l logs.Logger, s Storage) CreateGateway {
	return &CreateGtw{
		log:     l,
		Storage: s,
	}
}

