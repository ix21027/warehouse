package http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	controller Controller
	config     *Config
	router     *mux.Router
}

func NewServer(
	controller Controller,
) *Server {
	s := &Server{
		controller: controller,
		config:     NewConfig(),
		router:     mux.NewRouter(),
	}

	s.CreateRouter()

	return s
}

func (s *Server) Run() {
	go s.runHttpListener(s.config.Port)
	fmt.Println("Connected to HTTP servers")
	fmt.Println("HTTP servers start listening on port: ", s.config.Port)
}

func (s *Server) runHttpListener(port string) {
	fmt.Println("Connecting to HTTP...", port)
	if err := http.ListenAndServe(":"+port, s.router); err != nil {
		fmt.Println("Get error: ", err)
	}
}
