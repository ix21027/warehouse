package service

import (
	"log"
)

type GoodService struct{}

func NewGoodService() *GoodService {
	return &GoodService{}
}

type IGood interface {
	GetByID(string) string
	GetAll(string) string
}

func (s *GoodService) GetByID(data string) string {
	log.Println("get good by id service [data] : ", data)
	return ""
}

func (s *GoodService) GetAll(data string) string {
	log.Println("get all goods service [data] : ", data)
	return ""
}
