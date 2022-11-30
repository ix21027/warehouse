package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("create user handler")
}

//TODO: check the only account owner can update, delete and get user by id

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("update user handler [id] : ", mux.Vars(r)["id"])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("delete user handler [id] : ", mux.Vars(r)["id"])
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get user by id handler [id] : ", mux.Vars(r)["id"])
}
