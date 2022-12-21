package messaging

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
)

type Nats struct {
	Conn *nats.Conn
}

func New() *Nats {
	return &Nats{}
}

func Default() *Nats {
	n := New()
	n.Run()
	return n
}

func (n *Nats) Run() {
	if n.Conn == nil {
		n.Connect()
	}
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

//func (n *Nats) SendMsg(subj, msg string) {
//	if err := n.conn.Publish(subj, []byte(msg)); err != nil {
//		log.Println(err)
//	}
//}

//func (n *Nats) startSubscribers() {
//	_, err := n.conn.Subscribe("auth", MsgHandler)
//	if err != nil {
//		log.Println(err, " in startSubscribers()")
//	}
//}

//func MsgHandler(m *nats.Msg) {
//	//some business logic returns an answer that will be sent in response // server.svc.auth(m.Data)
//	err := m.Respond([]byte("answer is 42"))
//	if err != nil {
//		log.Println(err, " in MsgHandler()")
//	}
//}
