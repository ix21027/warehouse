package main

import (
	"customer_office_service/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func runRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Welcome!\n")); err != nil {
			log.Fatal(err)
		}
	})

	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(handler.CreateUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(handler.GetUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(handler.UpdateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(handler.DeleteUser)

	goodsR := r.PathPrefix("/goods").Subrouter()
	goodsR.Path("").Methods(http.MethodGet).HandlerFunc(handler.GetAllGoods)
	goodsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(handler.GetGoodByID)

	ordersR := r.PathPrefix("/orders").Subrouter()
	ordersR.Path("").Methods(http.MethodPost).HandlerFunc(handler.CreateOrder)
	ordersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(handler.UpdateOrder)
	ordersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(handler.GetOrderByID)
	ordersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(handler.DeleteOrder)

	log.Fatal(http.ListenAndServe(":8000", r))
}
