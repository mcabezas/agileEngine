package gateway

import (
	"context"
	"fmt"

	"github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type CreateGtw struct {
	accounts *gateway.AccountGateway
	log      logs.Logger
	Storage
}

func (g *CreateGtw) Create(ctx context.Context, cmd *models.CreateTransactionCMD) (*models.Transaction, error) {
	acc, err := g.accounts.Get(ctx, cmd.AccountID)
	if err != nil {
		err := fmt.Errorf("could not get account")
		g.log.Errorw("could not get account", logs.RequestID(ctx))
		return nil, err
	}
	initialBalance := acc.LockBalance()
	defer acc.UnlockBalance()
	var newBalance money.Money
	cmd.Amount = money.AmountAbs(cmd.Amount)
	if cmd.TransactionType == models.Credit {
		newBalance, err = initialBalance.Plus(cmd.Amount)
		if err != nil {
			g.log.Errorw("could not get account", logs.RequestID(ctx))
			return nil, err
		}
	}
	if cmd.TransactionType == models.Debit {
		newBalance, err = initialBalance.Minus(cmd.Amount)
		if err != nil {
			g.log.Errorw("could not get account", logs.RequestID(ctx))
			return nil, err
		}
		if newBalance.Amount < money.Amount(0) {
			g.log.Errorw("Negative Balance", logs.RequestID(ctx))
			return nil, fmt.Errorf("NegativeBalance")
		}
	}
	err = g.accounts.UpdateGateway.Update(ctx, acc, newBalance)
	if err != nil {
		return nil, err
	}
	transaction, err := g.Storage.Create(ctx, cmd)
	if err != nil {
		//Rollbacking balance
		if err = g.accounts.UpdateGateway.Update(ctx, acc, initialBalance); err != nil {
			g.log.Errorw("There was an issue rollbacking balance", logs.RequestID(ctx))
		}
	}
	return transaction, nil
}
