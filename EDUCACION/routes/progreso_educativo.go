package routes

import (
	"EDUCACION/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasProgresoEducativo(router *mux.Router) {
	router.HandleFunc("/progreso_educativo", controllers.ObtenerProgresosEducativos).Methods("GET")
	router.HandleFunc("/progreso_educativo/{id}", controllers.ObtenerProgresoEducativoPorID).Methods("GET")
	router.HandleFunc("/progreso_educativo", controllers.CrearProgresoEducativo).Methods("POST")
	router.HandleFunc("/progreso_educativo/{id}", controllers.ActualizarProgresoEducativo).Methods("PUT")
	router.HandleFunc("/progreso_educativo/{id}", controllers.EliminarProgresoEducativo).Methods("DELETE")
}
