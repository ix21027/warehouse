package nats_server

import (
	"log"
)

func (s *Server) SendToUserSvc(method string, payload []byte) {
	if err := s.Conn.Publish("userSvc."+method, payload); err != nil {
		log.Println(err)
	}
}
