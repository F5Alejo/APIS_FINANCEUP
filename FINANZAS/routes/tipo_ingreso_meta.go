package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTipoIngresoMeta(router *mux.Router) {
	router.HandleFunc("/tipo_ingreso_meta", controllers.ObtenerTiposIngresoMeta).Methods("GET")
	router.HandleFunc("/tipo_ingreso_meta/{id}", controllers.ObtenerTipoIngresoMetaPorID).Methods("GET")
	router.HandleFunc("/tipo_ingreso_meta", controllers.CrearTipoIngresoMeta).Methods("POST")
	router.HandleFunc("/tipo_ingreso_meta/{id}", controllers.ActualizarTipoIngresoMeta).Methods("PUT")
	router.HandleFunc("/tipo_ingreso_meta/{id}", controllers.EliminarTipoIngresoMeta).Methods("DELETE")
}
