package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasFinanzas(router *mux.Router) {
	router.HandleFunc("/finanzas", controllers.ObtenerRegistrosFinanzas).Methods("GET")
	router.HandleFunc("/finanzas/{id}", controllers.ObtenerFinanzasPorID).Methods("GET")
	router.HandleFunc("/finanzas", controllers.CrearFinanzas).Methods("POST")
	router.HandleFunc("/finanzas/{id}", controllers.ActualizarFinanzas).Methods("PUT")
	router.HandleFunc("/finanzas/{id}", controllers.EliminarFinanzas).Methods("DELETE")
}
