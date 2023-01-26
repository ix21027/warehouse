package nats_server

import (
	"github.com/nats-io/nats.go"
	"log"
)

func (s *Server) SendToUserSvc(method string, payload []byte) {
	m := nats.NewMsg("userSvc." + method)
	m.Header.Add("from", "customerOfficeSvc")
	m.Data = payload
	if err := s.Conn.PublishMsg(m); err != nil {
		log.Println("Error <SendToUserSvc>:", err)
	}

	//if err := s.NatsConn.Publish("userSvc."+method, payload); err != nil {
	//	log.Println(err)
	//}
}
