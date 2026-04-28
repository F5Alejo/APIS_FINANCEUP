package routes

import (
	"api_soporte/controller"

	"github.com/gorilla/mux"
)

// RegisterEstadoPqrRoutes registra las rutas para la tabla estado_pqr
func RegisterEstadoPqrRoutes(r *mux.Router) {
	r.HandleFunc("/estados-pqr", controller.GetAllEstadosPqr).Methods("GET")
	r.HandleFunc("/estados-pqr/{id}", controller.GetEstadoPqrByID).Methods("GET")
	r.HandleFunc("/estados-pqr", controller.CreateEstadoPqr).Methods("POST")
	r.HandleFunc("/estados-pqr/{id}", controller.UpdateEstadoPqr).Methods("PUT")
	r.HandleFunc("/estados-pqr/{id}", controller.DeleteEstadoPqr).Methods("DELETE")
}
