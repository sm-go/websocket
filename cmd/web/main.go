package main

import (
	"log"
	"net/http"

	"github.com/smith-golang/websocket/internal/handers"
)

func main() {
	mux := routes()

	log.Println("starting channel listener")
	go handers.ListenToWsChannel()

	log.Println("starting web server on port : 8080")
	_ = http.ListenAndServe(":8080", mux)
}
