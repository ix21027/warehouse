package service

type Service struct {
	Good  IGood
	Order IOrder
	User  IUser
}

type IGood interface {
	GetByID(any)
	GetAll(any)
}

type IOrder interface {
	Create(any)
	Update(any)
	Delete(any)
	GetByID(any)
}
type IUser interface {
	Create(any)
	Update(any)
	Delete(any)
	GetByID(any)
}

func New() *Service {
	return &Service{
		Good:  NewGoodService(),
		Order: NewOrderService(),
		User:  NewUserService(),
	}
}
