package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type ReqData struct {
	Code uint16 `json:"code"`
	Data any    `json:"data"`
}

func Reader(conn *websocket.Conn, ch chan string, s *Server) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Err reader: ", err)
			return
		}

		var data ReqData
		if err := json.Unmarshal(msg, &data); err != nil {
			fmt.Println("Err reader unmarshal msg: ", err)
			return
		}
		Route(data, ch, s)
	}
}
