package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *HttpController) CreateUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Create("Some business logic (data from CreateUser)")
	w.WriteHeader(201)
}

//TODO: check the only account owner can update, delete and get user by id

func (c *HttpController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Update("Some business logic (data from UpdateUser)")
	log.Println("update user handler [id] : ", mux.Vars(r)["id"])
}

func (c *HttpController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Delete("Some business logic (data from DeleteUser)")
	log.Println("delete user handler [id] : ", mux.Vars(r)["id"])
}

func (c *HttpController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	c.service.User.GetByID("Some business logic (data from GetUserByID)")
	log.Println("get user by id handler [id] : ", mux.Vars(r)["id"])
}

//  usersR := c.Router.PathPrefix("/users").Subrouter()
//	usersR.Path("").Methods(http.MethodPost).HandlerFunc(c.CreateUser)
//	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(c.GetUserByID)
//	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(c.UpdateUser)
//	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(c.DeleteUser)
