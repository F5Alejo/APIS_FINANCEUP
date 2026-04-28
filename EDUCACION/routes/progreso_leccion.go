package routes

import (
	"EDUCACION/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasProgresoLeccion(router *mux.Router) {
	router.HandleFunc("/progreso_leccion", controllers.ObtenerProgresosLeccion).Methods("GET")
	router.HandleFunc("/progreso_leccion/{id}", controllers.ObtenerProgresoLeccionPorID).Methods("GET")
	router.HandleFunc("/progreso_leccion", controllers.CrearProgresoLeccion).Methods("POST")
	router.HandleFunc("/progreso_leccion/{id}", controllers.ActualizarProgresoLeccion).Methods("PUT")
	router.HandleFunc("/progreso_leccion/{id}", controllers.EliminarProgresoLeccion).Methods("DELETE")
}
