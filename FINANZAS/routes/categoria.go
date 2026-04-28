package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasCategoria(router *mux.Router) {
	router.HandleFunc("/categoria", controllers.ObtenerCategorias).Methods("GET")
	router.HandleFunc("/categoria/{id}", controllers.ObtenerCategoriaPorID).Methods("GET")
	router.HandleFunc("/categoria", controllers.CrearCategoria).Methods("POST")
	router.HandleFunc("/categoria/{id}", controllers.ActualizarCategoria).Methods("PUT")
	router.HandleFunc("/categoria/{id}", controllers.EliminarCategoria).Methods("DELETE")
}
