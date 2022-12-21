package service

import (
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

func (s *UserService) Create(data any) {
	s.natsServer.SendMsg("authS.authorize", "Auth me please")
	log.Println("create user service [data] ", data)
}

//TODO: check the only account owner can update, delete and get user by id

func (s *UserService) Update(data any) {
	log.Println("update user service [data] : ", data)
}

func (s *UserService) Delete(data any) {
	log.Println("delete user service [data] : ", data)
}

func (s *UserService) GetByID(data any) {
	log.Println("get user by id service [data] : ", data)
}
