package server

import "net/http"

// TODO: research how to handle id from standard library
func (httpServer *Server) CreateRouter() {
	http.HandleFunc("/users", httpServer.controllers.CreateUser)
	//http.HandleFunc("/users/{id}", httpServer.controllers.UpdateUser)
	//http.HandleFunc("/users/{id}", httpServer.controllers.DeleteUser)
	//http.HandleFunc("/users/{id}", httpServer.controllers.GetUserByID)

	http.HandleFunc("/orders", httpServer.controllers.CreateOrder)
	//http.HandleFunc("/orders/{id}", httpServer.controllers.UpdateOrder)
	//http.HandleFunc("/orders/{id}", httpServer.controllers.DeleteOrder)
	//http.HandleFunc("/orders/{id}", httpServer.controllers.GetOrderByID)

	http.HandleFunc("/goods", httpServer.controllers.GetAllGoods)
	//http.HandleFunc("/goods/{id}", httpServer.controllers.GetGoodByID)
}
