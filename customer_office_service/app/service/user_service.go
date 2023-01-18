package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type UserService struct {
	natsServer INatsServer
}

type IUser interface {
	GetByStatus(string) string
	GetByLogin(string) string

	Create(string) string
	GetByID(string) string
	Update(string) string
	Delete(string) string
}

func NewUserService(ns INatsServer) *UserService {
	return &UserService{
		ns,
	}
}

func (s *UserService) Create(data string) string {
	createUserChanel := make(chan string)

	s.natsServer.Subscribe("customerOfficeSvc.CreateResp", func(m *nats.Msg) {
		createUserChanel <- string(m.Data)
	})
	s.natsServer.SendToUserSvc("Create", []byte(data))

	return <-createUserChanel
}

//TODO: check the only account owner can update, delete and get user by id

func (s *UserService) Update(data string) string {
	//TODO: implement
	payload, _ := json.Marshal(data)
	s.natsServer.SendToUserSvc("Update", payload)
	log.Println("update user service [data] : ", data)
	return ""
}

func (s *UserService) Delete(data string) string {
	deleteUserChanel := make(chan string)

	s.natsServer.Subscribe("customerOfficeSvc.DeleteResp", func(m *nats.Msg) {
		deleteUserChanel <- string(m.Data)
	})
	s.natsServer.SendToUserSvc("Delete", []byte(data))

	return <-deleteUserChanel
}

func (s *UserService) GetByID(data string) string {
	getByIDChanel := make(chan string)

	s.natsServer.Subscribe("customerOfficeSvc.GetByIDResp", func(m *nats.Msg) {
		getByIDChanel <- string(m.Data)
	})
	s.natsServer.SendToUserSvc("GetByID", []byte(data))

	return <-getByIDChanel
}
func (s *UserService) GetByStatus(data string) string {
	getByStatusChanel := make(chan string)

	s.natsServer.Subscribe("customerOfficeSvc.GetByStatusResp", func(m *nats.Msg) {
		getByStatusChanel <- string(m.Data)
	})
	s.natsServer.SendToUserSvc("GetByStatus", []byte(data))

	return <-getByStatusChanel
}
func (s *UserService) GetByLogin(data string) string {
	getByLoginChanel := make(chan string)

	//go func(c chan string) {
	//	s.natsServer.Subscribe("customerOfficeSvc.GetByLoginResp", func(m *nats.Msg) {
	//		c <- fmt.Sprintf("%s", m.Data)
	//	})
	//}(getByLoginChanel)

	s.natsServer.Subscribe("customerOfficeSvc.GetByLoginResp", func(m *nats.Msg) {
		getByLoginChanel <- string(m.Data)
	})
	s.natsServer.SendToUserSvc("GetByLogin", []byte(data))

	return <-getByLoginChanel
}
