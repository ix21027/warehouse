package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TODO: check only customer can create order

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	c.service.Order.Create("Some business logic (data from CreateOrder)")
	log.Println("create order handler")
	w.WriteHeader(http.StatusCreated)
}

//TODO: check that only order owner can update, delete and get order by id

func (c *Controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	c.service.Order.Update("Some business logic (data from UpdateOrder)")
	log.Println("update order handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	c.service.Order.Delete("Some business logic (data from DeleteOrder)")
	log.Println("delete order handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	c.service.Order.GetByID("Some business logic (data from GetOrderByID)")
	log.Println("get order by id handler [id] : ", mux.Vars(r)["id"])
}
