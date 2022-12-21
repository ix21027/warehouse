package service

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

type Service struct {
	nats *nats.Conn
}

func New(nc *nats.Conn) *Service {
	return &Service{nc}
}

func (s *Service) Authorize() {
	_, err := s.nats.Subscribe("authS.authorize", func(m *nats.Msg) {
		fmt.Println(string(m.Data), "gets from authS.Authorize")

		reply := "You have been successfully authorized"

		err := s.nats.Publish("customer_officeS.Authorize", []byte(reply))

		if err != nil {
			panic(err)
		}
	})

	if err != nil {
		panic(err)
	}
}
