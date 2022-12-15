package service

import (
	"net/http"
)

type Service struct {
	//Good  IGood
	//Order IOrder
	User IUser
}

//	type IGood interface {
//		GetByID(http.ResponseWriter, *http.Request)
//		GetAll(http.ResponseWriter, *http.Request)
//	}
//
//	type IOrder interface {
//		Create(http.ResponseWriter, *http.Request)
//		Update(http.ResponseWriter, *http.Request)
//		Delete(http.ResponseWriter, *http.Request)
//		GetByID(http.ResponseWriter, *http.Request)
//	}
type IUser interface {
	Create(data any)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
}

func New() *Service {
	return &Service{
		//Good:  NewGoodService(),
		//Order: NewOrderService(),
		User: NewUserService(),
	}
}
