package main

import (
	"admin_office_service/websocket"
)

func main() {
	s := websocket.RunServer()
	defer s.Stop()
}
