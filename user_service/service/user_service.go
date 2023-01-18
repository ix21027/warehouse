package service

import (
	"fmt"
	"strings"
	"user_service/redis"
	"user_service/scylla"
)

type SvcUser struct {
	rdb *redis.Redis
	sdb *scylla.Scylla
}
type UserService interface {
	UserInformer
}

func New(rdb *redis.Redis, scylla *scylla.Scylla) *SvcUser {
	return &SvcUser{rdb: rdb, sdb: scylla}
}

type UserInformer interface {
	UserInfo(string) (*redis.User, error)
	CreateUser(m string) string
	UpdateUser(m string) string
	DeleteUser(m string) string
	GetUserByID(m string) string
	GetUserByLogin(m string) string
	GetUserByStatus(m string) string
}

func (s *SvcUser) UserInfo(email string) (*redis.User, error) {
	u, err := s.rdb.GetUser(email)
	return u, err
}

func (s *SvcUser) CreateUser(m string) string {
	data := strings.Split(m, ",")
	name, login, password, typ := data[0], data[1], data[2], data[3]

	return s.sdb.InsertUser(name, login, password, typ)
}
func (s *SvcUser) UpdateUser(m string) string {
	data := strings.Split(m, ",")
	status, id := data[0], data[1]

	res := s.sdb.UpdateUserByStatus(status, id)
	if res != "" {
		return res
	}
	return "updated"
}
func (s *SvcUser) DeleteUser(m string) string {
	if err := s.sdb.DeleteUser(m); err != nil {
		return fmt.Sprintf("ERROR %s", err)
	}
	return "SUCCESSFULLY DELETED"
}
func (s *SvcUser) GetUserByID(m string) string {
	return s.sdb.GetUserByID(m)
}
func (s *SvcUser) GetUserByLogin(m string) string {
	return s.sdb.GetUserByLogin(m)
}
func (s *SvcUser) GetUserByStatus(m string) string {
	return s.sdb.GetUsersByStatus(m)
}
