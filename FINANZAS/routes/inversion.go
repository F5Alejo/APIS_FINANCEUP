package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasInversion(router *mux.Router) {
	router.HandleFunc("/inversion", controllers.ObtenerInversiones).Methods("GET")
	router.HandleFunc("/inversion/{id}", controllers.ObtenerInversionPorID).Methods("GET")
	router.HandleFunc("/inversion", controllers.CrearInversion).Methods("POST")
	router.HandleFunc("/inversion/{id}", controllers.ActualizarInversion).Methods("PUT")
	router.HandleFunc("/inversion/{id}", controllers.EliminarInversion).Methods("DELETE")
}
