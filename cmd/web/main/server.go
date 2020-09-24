package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/mcabezas/agileEngine/internal/logs"
)

type server struct {
	*http.Server
	log logs.Logger
}

func newServer(listening string, mux *chi.Mux, logger logs.Logger) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s, logger}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *server) Start() {
	srv.log.Info("starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.log.Fatalf("could not listen on %s due to %s", srv.Addr, err.Error())
		}
	}()
	srv.log.Infof("server is ready to handle requests %s", srv.Addr)

	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	srv.log.Infof("server is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.log.Fatalf("could not gracefully shutdown the server %s", err.Error())
	}

	srv.log.Info("server stopped")
}
