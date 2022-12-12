package websocket

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Server struct {
	NatsConn *nats.Conn
}

func New() *Server {
	nc, err := nats.Connect("nats", // or os.Getenv("NATS_URL")
		nats.MaxReconnects(3),
		nats.ReconnectWait(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{nc}
}

func (s *Server) SendNatsMsg(subj, msg string) error {
	if err := s.NatsConn.Publish(subj, []byte(msg)); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
