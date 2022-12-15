package controller_interface

import "net/http"

type Controller interface {
	IHttpController
}
type IHttpController interface {
	CreateUserController(w http.ResponseWriter, r *http.Request)
}
