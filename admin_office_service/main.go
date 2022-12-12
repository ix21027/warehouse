package main

import (
	server "admin_office_service/websocket"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	defer s.NatsConn.Close()

	http.HandleFunc("/ws", s.Endpoint)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
