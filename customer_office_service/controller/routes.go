package controller

import "net/http"

func (c *Controller) SetRouts() *Controller {
	usersR := c.Router.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(c.CreateUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(c.GetUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(c.UpdateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(c.DeleteUser)

	goodsR := c.Router.PathPrefix("/goods").Subrouter()
	goodsR.Path("").Methods(http.MethodGet).HandlerFunc(c.GetAllGoods)
	goodsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(c.GetGoodByID)

	ordersR := c.Router.PathPrefix("/orders").Subrouter()
	ordersR.Path("").Methods(http.MethodPost).HandlerFunc(c.CreateOrder)
	ordersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(c.UpdateOrder)
	ordersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(c.GetOrderByID)
	ordersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(c.DeleteOrder)

	return c
}
