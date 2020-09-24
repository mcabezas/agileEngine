package web

import (
	"github.com/go-chi/chi"
	"github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/commons/web"
)

type AccountHandler struct {
	*CreateHandler
	*GetHandler
}

func NewRoute(accountGateway *gateway.AccountGateway) web.Route {
	return &AccountHandler{
		NewCreateHandler(accountGateway.CreateGateway),
		NewGetHandler(accountGateway.GetGateway),
	}
}

func (h *AccountHandler) Up() *chi.Mux {
	router := chi.NewRouter()
	router.
		Group(func(r chi.Router) {
			r.Post("/", h.CreateHandler.Handle)
			r.Get("/", h.GetHandler.Handle)
		})
	return router
}
