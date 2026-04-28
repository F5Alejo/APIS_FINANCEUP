package routes

import (
	"EDUCACION/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasContenido(router *mux.Router) {
	router.HandleFunc("/contenido", controllers.ObtenerContenidos).Methods("GET")
	router.HandleFunc("/contenido/{id}", controllers.ObtenerContenidoPorID).Methods("GET")
	router.HandleFunc("/contenido", controllers.CrearContenido).Methods("POST")
	router.HandleFunc("/contenido/{id}", controllers.ActualizarContenido).Methods("PUT")
	router.HandleFunc("/contenido/{id}", controllers.EliminarContenido).Methods("DELETE")
}
