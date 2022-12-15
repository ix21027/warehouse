package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Order struct {
}

func NewOrderService() *Order {
	return &Order{}
}

//TODO: check only customer can create order

func (s *Order) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("create order handler")
	w.WriteHeader(http.StatusCreated)
}

//TODO: check that only order owner can update, delete and get order by id

func (s *Order) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("update order handler [id] : ", mux.Vars(r)["id"])
}

func (s *Order) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("delete order handler [id] : ", mux.Vars(r)["id"])
}

func (s *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get order by id handler [id] : ", mux.Vars(r)["id"])
}
