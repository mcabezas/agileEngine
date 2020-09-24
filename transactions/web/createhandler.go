package web

import (
	"encoding/json"
	"net/http"

	"github.com/mcabezas/agileEngine/internal/accounts"
	"github.com/mcabezas/agileEngine/internal/money"
	"github.com/mcabezas/agileEngine/transactions/gateway"
	"github.com/mcabezas/agileEngine/transactions/models"
)

type CreateHandler struct {
	gateway.CreateGateway
}

func NewCreateHandler(createGateway gateway.CreateGateway) *CreateHandler {
	return &CreateHandler{createGateway}
}

func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cmd, err := parseCreateRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.Create(ctx, cmd)
	if err != nil {
		//TODO It should not be always 400 but I have no time to make it better :(
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func parseCreateRequest(r *http.Request) (*models.CreateTransactionCMD, error) {
	body := r.Body
	defer func() {
		_ = body.Close()
	}()
	entity := &CreateTransactionCMD{}
	err := json.NewDecoder(body).Decode(&entity)

	if err != nil {
		return nil, err
	}
	return &models.CreateTransactionCMD{
		AccountID:       accounts.DefaultID,
		TransactionType: entity.TransactionType,
		Amount:          money.Money{
			Amount:   money.Amount(entity.Amount),
			Currency: money.DefaultCurrency(),
		},
	}, nil
}

type CreateTransactionCMD struct {
	TransactionType models.TransactionType `json:"type"`
	Amount          int64                  `json:"amount"`
}
