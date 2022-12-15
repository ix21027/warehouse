package service

import (
	"log"
)

type GoodService struct{}

func NewGoodService() *GoodService {
	return &GoodService{}
}

func (s *GoodService) GetByID(data any) {
	log.Println("get good by id service [data] : ", data)
}

func (s *GoodService) GetAll(data any) {
	log.Println("get all goods service [data] : ", data)
}
