package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasEditarMeta(router *mux.Router) {
	router.HandleFunc("/editar_meta", controllers.ObtenerEditarMetas).Methods("GET")
	router.HandleFunc("/editar_meta/{id}", controllers.ObtenerEditarMetaPorID).Methods("GET")
	router.HandleFunc("/editar_meta", controllers.CrearEditarMeta).Methods("POST")
	router.HandleFunc("/editar_meta/{id}", controllers.ActualizarEditarMeta).Methods("PUT")
	router.HandleFunc("/editar_meta/{id}", controllers.EliminarEditarMeta).Methods("DELETE")
}
