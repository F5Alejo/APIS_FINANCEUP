package routes

import (
	"EDUCACION/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasModuloEducativo(router *mux.Router) {
	router.HandleFunc("/modulo_educativo", controllers.ObtenerModulosEducativos).Methods("GET")
	router.HandleFunc("/modulo_educativo/{id}", controllers.ObtenerModuloEducativoPorID).Methods("GET")
	router.HandleFunc("/modulo_educativo", controllers.CrearModuloEducativo).Methods("POST")
	router.HandleFunc("/modulo_educativo/{id}", controllers.ActualizarModuloEducativo).Methods("PUT")
	router.HandleFunc("/modulo_educativo/{id}", controllers.EliminarModuloEducativo).Methods("DELETE")
}
