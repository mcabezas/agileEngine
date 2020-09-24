package gateway

import (
	"context"
	"testing"

	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateAccount(t *testing.T) {
	g := NewCreateGateway(logs.NewSugaredLogger(), newStorage())
	createdID, err := g.Create(context.Background())
	assert.True(t, createdID != "")
	assert.Nil(t, err)
}
