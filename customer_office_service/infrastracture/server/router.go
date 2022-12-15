package server

import "net/http"

func (httpServer *Server) CreateRouter() {
	http.HandleFunc("/users", httpServer.controllers.CreateUserController)
}
