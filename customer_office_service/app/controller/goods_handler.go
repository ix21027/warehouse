package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *Controller) GetGoodByID(w http.ResponseWriter, r *http.Request) {
	c.service.Good.GetByID("Some business logic (data from GetGoodByID)")
	log.Println("get good by id handler [id] : ", mux.Vars(r)["id"])
}

func (c *Controller) GetAllGoods(w http.ResponseWriter, r *http.Request) {
	c.service.Good.GetAll("Some business logic (data from GetAllGoods)")
	log.Println("get all goods handler")
	//w.Header().Set("Content-Code", "text/plain")
	w.WriteHeader(http.StatusOK)
}
