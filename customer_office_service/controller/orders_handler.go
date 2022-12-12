package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TODO: check only customer can create order

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("create order handler")
	w.WriteHeader(http.StatusCreated)
}

//TODO: check that only order owner can update, delete and get order by id

func (c *Controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("update order handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("delete order handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get order by id handler [id] : ", mux.Vars(r)["id"])
}
