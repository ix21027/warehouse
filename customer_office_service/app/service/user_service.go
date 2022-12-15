package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(data any) {
	//s.SendNatsMsg("test", "createGood")
	log.Println("create user service [data] ", data)
	//w.WriteHeader(200)
}

//TODO: check the only account owner can update, delete and get user by id

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("update user handler [id] : ", mux.Vars(r)["id"])
}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("delete user handler [id] : ", mux.Vars(r)["id"])
}

func (s *UserService) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get user by id handler [id] : ", mux.Vars(r)["id"])
}
