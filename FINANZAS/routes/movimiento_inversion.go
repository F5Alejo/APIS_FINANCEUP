package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasMovimientoInversion(router *mux.Router) {
	router.HandleFunc("/movimiento_inversion", controllers.ObtenerMovimientosInversion).Methods("GET")
	router.HandleFunc("/movimiento_inversion/{id}", controllers.ObtenerMovimientoInversionPorID).Methods("GET")
	router.HandleFunc("/movimiento_inversion", controllers.CrearMovimientoInversion).Methods("POST")
	router.HandleFunc("/movimiento_inversion/{id}", controllers.ActualizarMovimientoInversion).Methods("PUT")
	router.HandleFunc("/movimiento_inversion/{id}", controllers.EliminarMovimientoInversion).Methods("DELETE")
}
