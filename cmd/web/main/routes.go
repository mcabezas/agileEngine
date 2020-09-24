package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	accGtw "github.com/mcabezas/agileEngine/accounts/gateway"
	accWeb "github.com/mcabezas/agileEngine/accounts/web"
	trxGtw "github.com/mcabezas/agileEngine/transactions/gateway"
	trxWeb "github.com/mcabezas/agileEngine/transactions/web"
	"github.com/mcabezas/agileEngine/internal/logs"
)

func Routes(
	l logs.Logger) *chi.Mux {
	mux := chi.NewMux()
	basicCors := basicCors()

	//middlewares
	mux.Use(
		basicCors.Handler,
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
	)

	accStorage := accGtw.NewStorage()
	accountGateway := accGtw.NewGateway(accGtw.NewGetGateway(l,accStorage),
		accGtw.NewCreateGateway(l, accStorage),
		accGtw.NewUpdateGateway(l, accStorage))
	accountsRoute := accWeb.NewRoute(accountGateway)
	trxsRoute := trxWeb.NewRoute(l, trxGtw.NewStorage(), accountGateway)
	mux.Mount("/accounts", accountsRoute.Up())
	mux.Mount("/transactions", trxsRoute.Up())
	mux.Mount("/debug", middleware.Profiler())
	return mux
}

func basicCors() *cors.Cors {
	// Basic CORS
	// see: https://developer.github.com/v3/#cross-origin-resource-sharing
	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"}, //TODO load this configuration from a env variable
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	return c
}
