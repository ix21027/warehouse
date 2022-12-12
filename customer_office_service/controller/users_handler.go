package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("create user handler")
}

//TODO: check the only account owner can update, delete and get user by id

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("update user handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("delete user handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get user by id handler [id] : ", mux.Vars(r)["id"])
}
