package routes

import (
	"EDUCACION/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasLeccion(router *mux.Router) {
	router.HandleFunc("/leccion", controllers.ObtenerLecciones).Methods("GET")
	router.HandleFunc("/leccion/{id}", controllers.ObtenerLeccionPorID).Methods("GET")
	router.HandleFunc("/leccion", controllers.CrearLeccion).Methods("POST")
	router.HandleFunc("/leccion/{id}", controllers.ActualizarLeccion).Methods("PUT")
	router.HandleFunc("/leccion/{id}", controllers.EliminarLeccion).Methods("DELETE")
}
