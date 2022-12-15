package controller

import (
	"fmt"
	"net/http"
)

func (c *HttpController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- Som business logic 1---")
	c.service.User.Create("Some data")
	w.WriteHeader(201)
}
