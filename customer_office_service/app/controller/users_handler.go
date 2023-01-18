package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name     string `json:"name"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Type     string `json:"type"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := c.service.User.Create(u.Name + "," + u.Login + "," + u.Password + "," + u.Type)
	if res != "OK" {
		w.Write([]byte(res))
		return
	}
	w.WriteHeader(201)
}

//TODO: check the only account owner can update, delete and get user by id

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	c.service.User.Update(mux.Vars(r)["id"])
	w.Write([]byte(mux.Vars(r)["id"]))
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	res := c.service.User.Delete(mux.Vars(r)["id"])
	w.Write([]byte(res))
}

func (c *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	res := c.service.User.GetByID(mux.Vars(r)["id"])
	w.Write([]byte(res))
}
func (c *Controller) GetUserByLogin(w http.ResponseWriter, r *http.Request) {
	res := c.service.User.GetByLogin(mux.Vars(r)["login"])
	w.Write([]byte(res))
}
func (c *Controller) GetUserByStatus(w http.ResponseWriter, r *http.Request) {
	res := c.service.User.GetByStatus(mux.Vars(r)["status"])
	w.Write([]byte(res))
}
