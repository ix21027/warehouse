package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type ReqData struct {
	Code int               `json:"code"`
	Data map[string]string `json:"data"`
}

func Reader(conn *websocket.Conn, n *Server) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Err reader: ", err)
			return
		}

		var data ReqData
		if err := json.Unmarshal(msg, &data); err != nil {
			fmt.Println("Err reader unmarshal msga: ", err)
			return
		}
		Route(data, n)
	}
}
