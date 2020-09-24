package web

import (
	"github.com/go-chi/chi"
	"github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/commons/web"
)

type AccountHandler struct {
	*CreateHandler
}

func NewRoute(accountGateway *gateway.AccountGateway) web.Route {
	return &AccountHandler{NewCreateHandler(accountGateway.CreateGateway)}
}

func (h *AccountHandler) Up() *chi.Mux {
	router := chi.NewRouter()
	router.
		Group(func(r chi.Router) {
			r.Post("/", h.CreateHandler.Handle)
		})
	return router
}
