package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Create("Some business logic (data from CreateUser)")
	w.WriteHeader(201)
}

//TODO: check the only account owner can update, delete and get user by id

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Update("Some business logic (data from UpdateUser)")
	log.Println("update user handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Delete("Some business logic (data from DeleteUser)")
	log.Println("delete user handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	c.service.User.GetByID("Some business logic (data from GetUserByID)")
	log.Println("get user by id handler [id] : ", mux.Vars(r)["id"])
}
