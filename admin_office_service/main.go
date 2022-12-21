package main

import (
	"admin_office_service/websocket"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ws := websocket.New()
	defer ws.NatsConn.Close()

	http.HandleFunc("/ws", ws.Endpoint)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		return err
	}
	return nil
}
