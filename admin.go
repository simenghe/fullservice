package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.BasicAuth("BasicAuth", map[string]string{
		"Admin": "Password",
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Admin Router Reached"))
	})

	r.Get("/panel", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Panel Reached"))
	})
	return r
}
