package gateway

import (
	"context"

	"github.com/mcabezas/agileEngine/accounts/models"
	"github.com/mcabezas/agileEngine/internal"
	"github.com/mcabezas/agileEngine/internal/accounts"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
	"go.uber.org/zap"
)

type CreateGtw struct {
	log logs.Logger
	Storage
}

func (g *CreateGtw) Create(ctx context.Context) (string, error) {
	acc := models.NewAccount(internal.UUID(), money.Money{
		Amount: 0, Currency: money.DefaultCurrency(),
	})
	if err := g.Storage.Create(ctx, acc); err != nil {
		g.log.Errorw("There was an issue creating account", logs.RequestID(ctx))
		return "", err
	}
	g.log.Infow("Account created, from now it will be the default account", logs.RequestID(ctx), zap.String(logs.AccountID, acc.ID))
	accounts.DefaultID = acc.ID
	return acc.ID, nil
}
