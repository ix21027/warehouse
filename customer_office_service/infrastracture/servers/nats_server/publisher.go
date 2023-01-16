package nats_server

import (
	"log"
)

func (s *Server) SendToUserSvc(method string, payload []byte) {
	if err := s.conn.Publish("userSvc."+method+"User", payload); err != nil {
		log.Println(err)
	}
}
