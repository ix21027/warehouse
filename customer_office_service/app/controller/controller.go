package controller

import (
	"customer_office_service/app/service"
)

type HttpController struct {
	service *service.Service
}

func NewHTTPController(service *service.Service) *HttpController {
	return &HttpController{
		service,
	}
}
