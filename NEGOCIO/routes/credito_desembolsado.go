package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterCreditoDesembolsadoRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/creditos", controllers.GetAllCreditosDesembolsados).Methods("GET")
	r.HandleFunc("/negocio/creditos/{id}", controllers.GetCreditoDesembolsadoByID).Methods("GET")
	r.HandleFunc("/negocio/creditos", controllers.CreateCreditoDesembolsado).Methods("POST")
	r.HandleFunc("/negocio/creditos/{id}", controllers.UpdateCreditoDesembolsado).Methods("PUT")
	r.HandleFunc("/negocio/creditos/{id}", controllers.DeleteCreditoDesembolsado).Methods("DELETE")
}