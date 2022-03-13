package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	PORT = ":5000"
)

func Square(n int) int {
	return n * n
}

func main() {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	r.Get("/", HomePage)
	r.Mount("/admin", AdminRouter())

	log.Printf("Serving http server on port %s", PORT)
	http.ListenAndServe(PORT, r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page Reached\n"))
}
