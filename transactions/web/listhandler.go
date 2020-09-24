package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mcabezas/agileEngine/internal/accounts"
	"github.com/mcabezas/agileEngine/transactions/gateway"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type ListHandler struct {
	gateway.ListGateway
}

func NewListHandler(listGateway gateway.ListGateway) *ListHandler {
	return &ListHandler{listGateway}
}

func (h *ListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if accounts.DefaultID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ =w.Write([]byte("There is no account, it must be created first"))
		return
	}
	res, err := h.List(ctx, accounts.DefaultID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(mapResponse(res))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func mapResponse(trxs []*models.Transaction) []*TransactionResponse {
	if len(trxs) == 0 {
		return []*TransactionResponse{}
	}
	result := make([]*TransactionResponse, len(trxs))
	for i, t := range trxs {
		createdAt := time.Unix(t.CreatedAt, 0)
		result[i] = &TransactionResponse{
			ID:              t.ID,
			TransactionType: t.TransactionType,
			Amount:          int64(t.Amount.Amount),
			CreatedAt:       createdAt,
		}
	}
	return result
}

type TransactionResponse struct {
	ID              string
	TransactionType models.TransactionType
	Amount          int64
	CreatedAt       time.Time
}