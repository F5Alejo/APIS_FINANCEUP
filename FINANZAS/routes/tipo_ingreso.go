package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasTipoIngreso(router *mux.Router) {
	router.HandleFunc("/tipo_ingreso", controllers.ObtenerTiposIngreso).Methods("GET")
	router.HandleFunc("/tipo_ingreso/{id}", controllers.ObtenerTipoIngresoPorID).Methods("GET")
	router.HandleFunc("/tipo_ingreso", controllers.CrearTipoIngreso).Methods("POST")
	router.HandleFunc("/tipo_ingreso/{id}", controllers.ActualizarTipoIngreso).Methods("PUT")
	router.HandleFunc("/tipo_ingreso/{id}", controllers.EliminarTipoIngreso).Methods("DELETE")
}
