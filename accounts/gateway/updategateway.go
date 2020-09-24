package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/models"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
)

type UpdateGateway interface {
	Update(ctx context.Context, account *models.Account, newBalance money.Money) error
}

func NewUpdateGateway(l logs.Logger, s Storage) UpdateGateway {
	return &UpdateGtw{
		log:     l,
		Storage: s,
	}
}

