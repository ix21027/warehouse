package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
	"user_service/service"
)

type Nats struct {
	Conn *nats.Conn
	svc  service.UserService
}

func New(svc service.UserService) *Nats {
	return &Nats{svc: svc}
}

func Default(svc service.UserService) *Nats {
	n := New(svc)
	go n.Run()
	return n
}

func (n *Nats) Run() {
	if n.Conn == nil {
		n.Connect()
	}
	n.startSubscribers()
}

func (n *Nats) Connect() {
	fmt.Println("Connecting to NATS...")

	conn, err := n.connect()
	if err != nil {
		log.Println("err! Connect() :", err)
		return
	}
	n.Conn = conn

	fmt.Println("Connected to NATS")
}

func (n *Nats) connect() (*nats.Conn, error) {
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

func (n *Nats) Stop() {
	n.Conn.Close()
}

func (n *Nats) startSubscribers() {

	//Create
	if _, err := n.Conn.Subscribe("userSvc.Create", func(m *nats.Msg) {
		res := n.svc.CreateUser(fmt.Sprintf("%s", m.Data))
		n.Conn.Publish("customerOfficeSvc.CreateResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//GetUserByID
	if _, err := n.Conn.Subscribe("userSvc.GetByID", func(m *nats.Msg) {
		res := n.svc.GetUserByID(fmt.Sprintf("%s", m.Data))
		service := m.Header.Get("from")

		alowedServices := map[string]struct{}{"customerOfficeSvc": {}, "adminOfficeSvc": {}}
		if _, ok := alowedServices[service]; !ok {
			log.Println("there is no such service in allowed")
			return
		}
		n.SendTo(service, "GetByIDResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//GetUserByStatus
	if _, err := n.Conn.Subscribe("userSvc.GetByStatus", func(m *nats.Msg) {
		res := n.svc.GetUserByStatus(fmt.Sprintf("%s", m.Data))
		n.SendTo(m.Header.Get("from"), "GetByStatusResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//GetUserByLogin
	if _, err := n.Conn.Subscribe("userSvc.GetByLogin", func(m *nats.Msg) {
		res := n.svc.GetUserByLogin(fmt.Sprintf("%s", m.Data))
		n.SendTo(m.Header.Get("from"), "GetByLoginResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//DeleteUser
	if _, err := n.Conn.Subscribe("userSvc.Delete", func(m *nats.Msg) {
		res := n.svc.DeleteUser(fmt.Sprintf("%s", m.Data))
		n.SendTo(m.Header.Get("from"), "DeleteResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//UpdateUser (not implemented on other service)
	if _, err := n.Conn.Subscribe("userSvc.Update", func(m *nats.Msg) {
		res := n.svc.UpdateUser(fmt.Sprintf("%s", m.Data))
		n.Conn.Publish("customerOfficeSvc.UpdateResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}

	//***** GET REQ FROM ADMIN OFFICE SERVICE *****//
	if _, err := n.Conn.Subscribe("userSvc.Update", func(m *nats.Msg) {
		res := n.svc.UpdateUser(fmt.Sprintf("%s", m.Data))
		n.Conn.Publish("customerOfficeSvc.UpdateResp", []byte(res))
	}); err != nil {
		log.Println(err, "ERROR in startSubscribers()")
	}
}

func (n *Nats) SendTo(svc, method string, payload []byte) {
	m := nats.NewMsg(svc + "." + method)
	m.Header.Add("from", "userSvc")
	m.Header.Add("ResponseCode", method)
	m.Data = payload
	if err := n.Conn.PublishMsg(m); err != nil {
		log.Println("Error <SendTo>:", err)
	}
}
