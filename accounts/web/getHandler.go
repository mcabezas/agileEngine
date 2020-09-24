package web

import (
	"encoding/json"
	"net/http"

	"github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/accounts"
	"github.com/mcabezas/agileEngine/internal/money"
)

type GetHandler struct {
	gateway.GetGateway
}

func NewGetHandler(getGateway gateway.GetGateway) *GetHandler {
	return &GetHandler{getGateway}
}

func (h *GetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if accounts.DefaultID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("There is no account, it must be created first"))
		return
	}
	res, err := h.Get(ctx, accounts.DefaultID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(&AccountResponse{
		ID:      res.ID,
		Balance: res.Balance(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

type AccountResponse struct {
	ID      string      `json:"id"`
	Balance money.Money `json:"balance"`
}
