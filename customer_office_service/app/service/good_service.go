package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type GoodService struct{}

func NewGoodService() *GoodService {
	return &GoodService{}
}

func (gs *GoodService) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get good by id handler [id] : ", mux.Vars(r)["id"])
}

func (gs *GoodService) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("get all goods handler")
	//w.Header().Set("Content-Code", "text/plain")
	w.WriteHeader(http.StatusOK)
}
