package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetGoodByID(w http.ResponseWriter, r *http.Request) {
	log.Println("get good by id handler [id] : ", mux.Vars(r)["id"])
}

func GetAllGoods(w http.ResponseWriter, r *http.Request) {
	log.Println("get all goods handler")
	//w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
