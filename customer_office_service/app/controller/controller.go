package controller

import (
	"customer_office_service/app/service"
)

type Controller struct {
	service *service.Service
}

func New(service *service.Service) *Controller {
	return &Controller{
		service,
	}
}
