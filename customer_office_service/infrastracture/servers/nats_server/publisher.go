package nats_server

import (
	"log"
)

func (s *Server) SendMsg(subj, msg string) {
	if err := s.conn.Publish(subj, []byte(msg)); err != nil {
		log.Println(err)
	}
}
