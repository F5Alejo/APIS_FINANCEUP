package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasNivelRiesgo(router *mux.Router) {
	router.HandleFunc("/nivel_riesgo", controllers.ObtenerNivelesRiesgo).Methods("GET")
	router.HandleFunc("/nivel_riesgo/{id}", controllers.ObtenerNivelRiesgoPorID).Methods("GET")
	router.HandleFunc("/nivel_riesgo", controllers.CrearNivelRiesgo).Methods("POST")
	router.HandleFunc("/nivel_riesgo/{id}", controllers.ActualizarNivelRiesgo).Methods("PUT")
	router.HandleFunc("/nivel_riesgo/{id}", controllers.EliminarNivelRiesgo).Methods("DELETE")
}
