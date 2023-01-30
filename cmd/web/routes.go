package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/smith-golang/websocket/internal/handers"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handers.Home))
	mux.Get("/ws", http.HandlerFunc(handers.WsEndPoint))
	return mux
}
