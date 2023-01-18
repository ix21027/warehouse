package service

import "github.com/nats-io/nats.go"

type Service struct {
	Good       IGood
	Order      IOrder
	User       IUser
	natsServer INatsServer
}

func NewMain(natsServer INatsServer) *Service {
	return &Service{
		Good:       NewGoodService(),
		Order:      NewOrderService(),
		User:       NewUserService(natsServer),
		natsServer: natsServer,
	}
}

type INatsServer interface {
	SendToUserSvc(method string, payload []byte)
	Subscribe(subj string, cb nats.MsgHandler)
}
