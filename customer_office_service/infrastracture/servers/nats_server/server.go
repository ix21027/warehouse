package nats_server

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Server struct {
	conn   *nats.Conn
	config *Config
}

func New() *Server {
	return &Server{
		config: NewConfig(),
	}
}

func (s *Server) Run() {
	if s.conn == nil {
		s.Connect()
	}
	//s.startSubscribers()
}

func (s *Server) Connect() {
	fmt.Println("Connecting to NATS...")

	conn, err := s.connect()
	if err != nil {
		log.Println("err! Connect() :", err)
		return
	}
	s.conn = conn

	fmt.Println("Connected to NATS")
}

func (s *Server) connect() (*nats.Conn, error) {
	return nats.Connect(
		s.config.Address,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(5*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Println("Disconnected from NATS")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Println("Reconnecting to NATS...")
		}),
	)
}

func (s *Server) Stop() {
	s.conn.Close()
}
