package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasMovimientoMeta(router *mux.Router) {
	router.HandleFunc("/movimiento_meta", controllers.ObtenerMovimientosMeta).Methods("GET")
	router.HandleFunc("/movimiento_meta/{id}", controllers.ObtenerMovimientoMetaPorID).Methods("GET")
	router.HandleFunc("/movimiento_meta", controllers.CrearMovimientoMeta).Methods("POST")
	router.HandleFunc("/movimiento_meta/{id}", controllers.ActualizarMovimientoMeta).Methods("PUT")
	router.HandleFunc("/movimiento_meta/{id}", controllers.EliminarMovimientoMeta).Methods("DELETE")
}
