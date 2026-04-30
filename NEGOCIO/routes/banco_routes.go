package routes

import (
	"NEGOCIO/controllers"

	"github.com/gorilla/mux"
)

func RegisterBancoRoutes(r *mux.Router) {
	r.HandleFunc("/bancos", controllers.GetAllBancos).Methods("GET")
	r.HandleFunc("/bancos/{id}", controllers.GetBancoByID).Methods("GET")
	r.HandleFunc("/bancos", controllers.CreateBanco).Methods("POST")
	r.HandleFunc("/bancos/{id}", controllers.UpdateBanco).Methods("PUT")
	r.HandleFunc("/bancos/{id}", controllers.DeleteBanco).Methods("DELETE")
}