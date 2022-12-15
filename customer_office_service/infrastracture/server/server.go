package server

import (
	"customer_office_service/infrastracture/server/interfaces"
	"fmt"
	"net/http"
)

type Server struct {
	controllers controller_interface.Controller
	config      *Config
}

func NewHTTPServer(
	conf *Config,
	controllers controller_interface.Controller,
) (*Server, error) {
	httpServer := &Server{
		controllers: controllers,
		config:      conf,
	}

	httpServer.CreateRouter()

	return httpServer, nil
}

func (httpServer *Server) Run() {
	go httpServer.runHttpListener(httpServer.config.Port)
	fmt.Println("Connected to HTTP server")
	fmt.Println("HTTP server start listening on port: ", httpServer.config.Port)
}

func (httpServer *Server) runHttpListener(port string) {
	fmt.Println("Connecting to HTTP...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Get error: ", err)
	}
}
