package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/models"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
)

type UpdateGtw struct {
	log logs.Logger
	Storage
}

func (g *UpdateGtw) Update(ctx context.Context, account *models.Account, newBalance money.Money) error {
	acc := models.NewAccount(account.ID, newBalance)
	if err := g.Storage.Update(ctx, acc); err == nil {
		g.log.Errorw("There was an issue updating account", logs.RequestID(ctx))
		return err
	}
	return nil
}
