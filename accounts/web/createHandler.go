package web

import (
	"encoding/json"
	"net/http"

	"github.com/mcabezas/agileEngine/accounts/gateway"
)

type CreateHandler struct {
	gateway.CreateGateway
}

func NewCreateHandler(createGateway gateway.CreateGateway) *CreateHandler {
	return &CreateHandler{createGateway}
}

func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.Create(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

