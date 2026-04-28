package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTipoIngresoInversion(router *mux.Router) {
	router.HandleFunc("/tipo_ingreso_inversion", controllers.ObtenerTiposIngresoInversion).Methods("GET")
	router.HandleFunc("/tipo_ingreso_inversion/{id}", controllers.ObtenerTipoIngresoInversionPorID).Methods("GET")
	router.HandleFunc("/tipo_ingreso_inversion", controllers.CrearTipoIngresoInversion).Methods("POST")
	router.HandleFunc("/tipo_ingreso_inversion/{id}", controllers.ActualizarTipoIngresoInversion).Methods("PUT")
	router.HandleFunc("/tipo_ingreso_inversion/{id}", controllers.EliminarTipoIngresoInversion).Methods("DELETE")
}
