package service

import (
	"log"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

//TODO: check only customer can create order

func (s *OrderService) Create(data any) {
	log.Println("create order service [data]")
}

//TODO: check that only order owner can update, delete and get order by id

func (s *OrderService) Update(data any) {
	log.Println("update order service [data] : ", data)
}

func (s *OrderService) Delete(data any) {
	log.Println("delete order service [data] : ", data)
}

func (s *OrderService) GetByID(data any) {
	log.Println("get order by id service [data] : ", data)
}
