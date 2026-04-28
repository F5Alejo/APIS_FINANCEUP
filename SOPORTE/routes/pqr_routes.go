package routes

import (
	"soporte/controller"

	"github.com/gorilla/mux"
)

// RegisterPqrRoutes registra las rutas para la tabla pqr
func RegisterPqrRoutes(r *mux.Router) {
	r.HandleFunc("/pqrs", controller.GetAllPqrs).Methods("GET")
	r.HandleFunc("/pqrs/{id}", controller.GetPqrByID).Methods("GET")
	r.HandleFunc("/pqrs", controller.CreatePqr).Methods("POST")
	r.HandleFunc("/pqrs/{id}", controller.UpdatePqr).Methods("PUT")
	r.HandleFunc("/pqrs/{id}", controller.DeletePqr).Methods("DELETE")
}
