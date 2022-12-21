package service

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
