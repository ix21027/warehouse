package service

import (
	"log"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

type IOrder interface {
	Create(string) string
	GetByID(string) string
	Update(string) string
	Delete(string) string
}

//TODO: check only customer can create order

func (s *OrderService) Create(data string) string {
	log.Println("create order service [data]")
	return ""
}

//TODO: check that only order owner can update, delete and get order by id

func (s *OrderService) Update(data string) string {
	log.Println("update order service [data] : ", data)
	return ""
}

func (s *OrderService) Delete(data string) string {
	log.Println("delete order service [data] : ", data)
	return ""
}

func (s *OrderService) GetByID(data string) string {
	log.Println("get order by id service [data] : ", data)
	return ""
}
