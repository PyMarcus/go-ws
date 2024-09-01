package main

import (
	"net/http"

	"github.com/PyMarcus/go_websockets/internal"
	"github.com/bmizerany/pat"
)

// routes the application
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(internal.Home))
	mux.Get("/ws", http.HandlerFunc(internal.WsEndpoint))

	return mux
}
