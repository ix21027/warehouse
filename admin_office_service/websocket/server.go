package websocket

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	NatsConn *nats.Conn
}

func RunServer() *Server {
	s := &Server{}
	s.Run()
	return s
}

func (s *Server) Run() {
	if s.NatsConn == nil {
		s.Connect()
	}
	s.startSubscribers()

	http.HandleFunc("/ws", s.Endpoint)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	s.NatsConn.Close()
}

func (s *Server) Connect() {
	fmt.Println("Connecting to NATS...")

	conn, err := s.connect()
	if err != nil {
		log.Println("err! Connect() :", err)
		return
	}
	s.NatsConn = conn

	fmt.Println("Connected to NATS")
}

func (s *Server) connect() (*nats.Conn, error) {
	return nats.Connect(
		os.Getenv("NATS_URL"),
		nats.MaxReconnects(2),
		nats.ReconnectWait(5*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Println("Disconnected from NATS")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Println("Reconnecting to NATS...")
		}),
	)
}

func (s *Server) SendNatsMsg(subj, msg string) error {
	if err := s.NatsConn.Publish(subj, []byte(msg)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *Server) SendToUserSvc(method string, payload []byte) {
	m := nats.NewMsg("userSvc." + method)
	m.Header.Add("from", "adminOfficeSvc")
	m.Data = payload
	if err := s.NatsConn.PublishMsg(m); err != nil {
		log.Println(err)
	}
	log.Println("subj:", "'userSvc."+method+"'", "data:", string(payload))
}

func (s *Server) SendReqToUserSvcGetByID(id []byte) {
	s.SendToUserSvc("GetByID", id)
}
func (s *Server) SendReqToUserSvcGetByStatus(status []byte) {
	s.SendToUserSvc("GetByStatus", status)
}
func (s *Server) SendReqToUserSvcGetByLogin(login []byte) {
	s.SendToUserSvc("GetByLogin", login)
}

func (s *Server) startSubscribers() {
	if _, err := s.NatsConn.Subscribe("adminOfficeSvc.*", func(m *nats.Msg) {
		switch m.Header.Get("ResponseCode") {
		case "GetByIDResp":
			ch <- string(m.Data)
		case "GetByStatusResp":
			ch <- string(m.Data)
		case "GetByLoginResp":
			ch <- string(m.Data)
		default:
			log.Println("Unknown response from nats server")
		}

	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}
}
