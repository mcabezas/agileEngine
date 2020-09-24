package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/mcabezas/agileEngine/internal/accounts"
	"github.com/mcabezas/agileEngine/transactions/gateway"
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
		_, _ =w.Write([]byte("There is no account, it must be created first"))
		return
	}
	transactionID := chi.URLParam(r, "transactionId")
	res, err := h.Get(ctx, transactionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	createdAt := time.Unix(res.CreatedAt, 0)
	t := &TransactionResponse{
		ID:              res.ID,
		TransactionType: res.TransactionType,
		Amount:          int64(res.Amount.Amount),
		CreatedAt:       createdAt,
	}
	result, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}