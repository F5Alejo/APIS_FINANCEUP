package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTipoInversion(router *mux.Router) {
	router.HandleFunc("/tipo_inversion", controllers.ObtenerTiposInversion).Methods("GET")
	router.HandleFunc("/tipo_inversion/{id}", controllers.ObtenerTipoInversionPorID).Methods("GET")
	router.HandleFunc("/tipo_inversion", controllers.CrearTipoInversion).Methods("POST")
	router.HandleFunc("/tipo_inversion/{id}", controllers.ActualizarTipoInversion).Methods("PUT")
	router.HandleFunc("/tipo_inversion/{id}", controllers.EliminarTipoInversion).Methods("DELETE")
}
