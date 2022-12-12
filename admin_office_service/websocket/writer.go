package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func Writer(conn *websocket.Conn, pongCh chan byte, ch chan string) {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		//case <-u.Ctx.Done():
		//	return
		case message, ok := <-ch:
			if !ok {
				return
			}

			if err := conn.SetWriteDeadline(time.Now().Add(15 * time.Second)); err != nil {
				fmt.Println("!!! Get err SetWriteDeadline: ", err)
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Println("!!! Get err w.Write([]byte(message)): ", err)
				return
			}

		case _, ok := <-pongCh:
			if !ok {
				return
			}
			if err := conn.SetWriteDeadline(time.Now().Add(15 * time.Second)); err != nil {
				return
			}
			err := conn.WriteMessage(websocket.PongMessage, nil)
			if err != nil {
				fmt.Println("!!! Get err writer <-pongCh: ", err)
				return
			}

		case <-ticker.C:
			if err := conn.SetWriteDeadline(time.Now().Add(15 * time.Second)); err != nil {
				fmt.Println("!!! Get err: ", err)
				return
			}

			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println("!!! Get err: ", err)
				return
			}
		}
	}
}
