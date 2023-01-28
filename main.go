package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade : ", err)
		}
		defer conn.Close()
		for i := 0; i < 30; i++ {
			time.Sleep(time.Second * 1)
			conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
		}
	})
	fmt.Println("Server is starting on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
