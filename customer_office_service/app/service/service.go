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

type IGood interface {
	GetByID(any)
	GetAll(any)
}

type IOrder interface {
	CRUDer
}
type IUser interface {
	CRUDer
}

type CRUDer interface {
	Create(any)
	GetByID(any)
	Update(any)
	Delete(any)
}

type INatsServer interface {
	SendToUserSvc(method string, payload []byte)
}
