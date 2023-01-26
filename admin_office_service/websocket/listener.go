package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

func (s *Server) RunListener(conn *websocket.Conn, ch chan string) error {
	pongCh := make(chan byte)

	if err := setUpConnHandler(conn, pongCh); err != nil {
		return err
	}

	//ch := make(chan string)

	go Reader(conn, s)
	go Writer(conn, pongCh, ch)

	for {
	}

	return nil
}

func setUpConnHandler(conn *websocket.Conn, pongCh chan byte) error {
	readDeadLine := 15 * time.Second
	conn.SetCloseHandler(func(_ int, _ string) error {
		err := conn.Close()
		if err != nil {
			return err
		}
		return nil
	})
	conn.SetReadLimit(1024)

	if err := conn.SetReadDeadline(time.Now().Add(readDeadLine)); err != nil {
		return err
	}

	conn.SetPongHandler(
		func(string) error {
			return conn.SetReadDeadline(time.Now().Add(readDeadLine))
		},
	)

	conn.SetPingHandler(
		func(string) error {
			pongCh <- 0
			return conn.SetReadDeadline(time.Now().Add(readDeadLine))
		},
	)
	return nil
}
