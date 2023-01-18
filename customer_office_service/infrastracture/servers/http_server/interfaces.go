package http_server

import "net/http"

type Controller interface {
	IUser
	IOrder
	IGood
}

type IUser interface {
	CreateUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
	GetUserByLogin(http.ResponseWriter, *http.Request)
	GetUserByStatus(http.ResponseWriter, *http.Request)
}

type IOrder interface {
	CreateOrder(http.ResponseWriter, *http.Request)
	UpdateOrder(http.ResponseWriter, *http.Request)
	DeleteOrder(http.ResponseWriter, *http.Request)
	GetOrderByID(http.ResponseWriter, *http.Request)
}

type IGood interface {
	GetGoodByID(http.ResponseWriter, *http.Request)
	GetAllGoods(http.ResponseWriter, *http.Request)
}
