package web

import (
	"github.com/go-chi/chi"
	accGtw "github.com/mcabezas/agileEngine/accounts/gateway"
	"github.com/mcabezas/agileEngine/internal/commons/web"
	"github.com/mcabezas/agileEngine/internal/logs"
	"github.com/mcabezas/agileEngine/transactions/gateway"
)

type TransactionHandler struct {
	*CreateHandler
	*ListHandler
	*GetHandler
}

func NewRoute(l logs.Logger, storage gateway.Storage, accounts *accGtw.AccountGateway) web.Route {
	return &TransactionHandler{
		CreateHandler: NewCreateHandler(gateway.NewCreateGateway(l, storage, accounts)),
		ListHandler: NewListHandler(gateway.NewListGateway(l, storage)),
		GetHandler: NewGetHandler(gateway.NewGetGateway(l, storage)),
	}
}

func (h *TransactionHandler) Up() *chi.Mux {
	router := chi.NewRouter()
	router.
		Group(func(r chi.Router) {
			r.Post("/", h.CreateHandler.Handle)
			r.Get("/{transactionId}", h.GetHandler.Handle)
			r.Get("/history", h.ListHandler.Handle)
		})
	return router
}

