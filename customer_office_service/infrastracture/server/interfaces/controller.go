package controller_interface

import "net/http"

type Controller interface {
	IHttpController
}
type IHttpController interface {
	CreateUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)

	CreateOrder(http.ResponseWriter, *http.Request)
	UpdateOrder(http.ResponseWriter, *http.Request)
	DeleteOrder(http.ResponseWriter, *http.Request)
	GetOrderByID(http.ResponseWriter, *http.Request)

	GetGoodByID(http.ResponseWriter, *http.Request)
	GetAllGoods(http.ResponseWriter, *http.Request)
}
