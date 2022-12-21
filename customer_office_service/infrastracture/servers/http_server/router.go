package http_server

import (
	"net/http"
)

func (s *Server) CreateRouter() {
	usersR := s.router.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(s.controller.CreateUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(s.controller.GetUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(s.controller.UpdateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(s.controller.DeleteUser)

	ordersR := s.router.PathPrefix("/orders").Subrouter()
	ordersR.Path("").Methods(http.MethodPost).HandlerFunc(s.controller.CreateOrder)
	ordersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(s.controller.UpdateOrder)
	ordersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(s.controller.GetOrderByID)
	ordersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(s.controller.DeleteOrder)

	goodsR := s.router.PathPrefix("/goods").Subrouter()
	goodsR.Path("").Methods(http.MethodGet).HandlerFunc(s.controller.GetAllGoods)
	goodsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(s.controller.GetGoodByID)
}
