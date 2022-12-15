package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *HttpController) GetGoodByID(w http.ResponseWriter, r *http.Request) {
	c.service.Good.GetByID("Some business logic (data from GetGoodByID)")
	log.Println("get good by id handler [id] : ", mux.Vars(r)["id"])
}

func (c *HttpController) GetAllGoods(w http.ResponseWriter, r *http.Request) {
	c.service.Good.GetAll("Some business logic (data from GetAllGoods)")
	log.Println("get all goods handler")
	//w.Header().Set("Content-Code", "text/plain")
	w.WriteHeader(http.StatusOK)
}

//func (c *HttpController) SetGoodsRouts() {
//	goodsR := c.Router.PathPrefix("/goods").Subrouter()
//	goodsR.Path("").Methods(http.MethodGet).HandlerFunc(c.service.Good.GetAll)
//	goodsR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(c.service.Good.GetByID)
//}
