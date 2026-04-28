package routes

import (
	"EDUCACION/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/modulo_educativo", handlers.GetModulosEducativos).Methods("GET")
	router.HandleFunc("/modulo_educativo/{id}", handlers.GetModuloEducativoByID).Methods("GET")
	router.HandleFunc("/modulo_educativo", handlers.CreateModuloEducativo).Methods("POST")
	router.HandleFunc("/modulo_educativo/{id}", handlers.UpdateModuloEducativo).Methods("PUT")
	router.HandleFunc("/modulo_educativo/{id}", handlers.DeleteModuloEducativo).Methods("DELETE")

	return router
}
