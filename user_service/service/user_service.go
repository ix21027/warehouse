package service

import "user_service/redis"

type SUser struct {
	rdb *redis.Redis
}
type UserService interface {
	UserInformer
}

func New(rdb *redis.Redis) *SUser {
	return &SUser{rdb: rdb}
}

type UserInformer interface {
	UserInfo(string) (*redis.User, error)
}

func (s *SUser) UserInfo(email string) (*redis.User, error) {
	u, err := s.rdb.GetUser(email)

	return u, err
}
