package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller struct {
	Router *mux.Router
}

func New() *Controller {
	return &Controller{
		mux.NewRouter(),
	}
}

func (c *Controller) ServeHTTP() {
	log.Fatal(http.ListenAndServe(":8000", c.Router))
}
