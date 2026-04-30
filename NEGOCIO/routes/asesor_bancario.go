package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterAsesorRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/asesores", controllers.GetAllAsesores).Methods("GET")
	r.HandleFunc("/negocio/asesores/{id}", controllers.GetAsesorByID).Methods("GET")
	r.HandleFunc("/negocio/asesores", controllers.CreateAsesor).Methods("POST")
	r.HandleFunc("/negocio/asesores/{id}", controllers.UpdateAsesor).Methods("PUT")
	r.HandleFunc("/negocio/asesores/{id}", controllers.DeleteAsesor).Methods("DELETE")
}