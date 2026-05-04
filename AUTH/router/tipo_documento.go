package router

import (
	"AUTH/controller"

	"github.com/gorilla/mux"
)

func RegisterTipoDocumentoRoutes(r *mux.Router) {
	r.HandleFunc("/auth/tipos-documento", controller.GetAllTiposDocumento).Methods("GET")
	r.HandleFunc("/auth/tipos-documento/{id}", controller.GetTipoDocumentoByID).Methods("GET")
	r.HandleFunc("/auth/tipos-documento", controller.CreateTipoDocumento).Methods("POST")
	r.HandleFunc("/auth/tipos-documento/{id}", controller.UpdateTipoDocumento).Methods("PUT")
	r.HandleFunc("/auth/tipos-documento/{id}", controller.DeleteTipoDocumento).Methods("DELETE")
}
