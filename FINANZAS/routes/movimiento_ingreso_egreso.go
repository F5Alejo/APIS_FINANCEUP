package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasMovimientoIngresoEgreso(router *mux.Router) {
	router.HandleFunc("/movimiento_ingreso_egreso", controllers.ObtenerMovimientosIngresoEgreso).Methods("GET")
	router.HandleFunc("/movimiento_ingreso_egreso/{id}", controllers.ObtenerMovimientoIngresoEgresoPorID).Methods("GET")
	router.HandleFunc("/movimiento_ingreso_egreso", controllers.CrearMovimientoIngresoEgreso).Methods("POST")
	router.HandleFunc("/movimiento_ingreso_egreso/{id}", controllers.ActualizarMovimientoIngresoEgreso).Methods("PUT")
	router.HandleFunc("/movimiento_ingreso_egreso/{id}", controllers.EliminarMovimientoIngresoEgreso).Methods("DELETE")
}
