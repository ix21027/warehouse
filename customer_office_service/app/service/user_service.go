package service

import (
	"encoding/json"
	"log"
)

type UserService struct {
	natsServer INatsServer
}

func NewUserService(ns INatsServer) *UserService {
	return &UserService{
		ns,
	}
}

type User struct {
	Email        string
	PasswordHash string
}

func (s *UserService) Create(data any) {
	user, ok := data.(User)
	if !ok {
		log.Println("Err: data is not a User!")
		return
	}
	payload, _ := json.Marshal(user)
	s.natsServer.SendToUserSvc("Create", payload)
	log.Println("create user service [data] ", data)
}

//TODO: check the only account owner can update, delete and get user by id

func (s *UserService) Update(data any) {
	payload, _ := json.Marshal(data)
	s.natsServer.SendToUserSvc("Update", payload)
	log.Println("update user service [data] : ", data)
}

func (s *UserService) Delete(data any) {
	payload, _ := json.Marshal(data)
	s.natsServer.SendToUserSvc("Delete", payload)
	log.Println("delete user service [data] : ", data)
}

func (s *UserService) GetByID(data any) {
	payload, _ := json.Marshal(data)
	s.natsServer.SendToUserSvc("GetByID", payload)
	log.Println("get user by id service [data] : ", data)
}
