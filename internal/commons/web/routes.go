package web

import "github.com/go-chi/chi"

type Route interface {
	Up() *chi.Mux
}
