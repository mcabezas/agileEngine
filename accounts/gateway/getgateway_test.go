package gateway

import (
	"context"
	"testing"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/internal/money"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateAndGet(t *testing.T) {
	storage := newStorage()
	createGateway := NewCreateGateway(logs.NewSugaredLogger(), storage)
	createdID, err := createGateway.Create(context.Background())
	getGateway := NewGetGateway(logs.NewSugaredLogger(), storage)
	foundAcc, err := getGateway.Get(context.Background(), createdID)
	assert.Nil(t, err)
	assert.Equal(t, createdID, foundAcc.ID)
	assert.Equal(t, foundAcc.Balance.Amount, money.Amount(0))
}
