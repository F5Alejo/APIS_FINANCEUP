package routes

import (
	"EDUCACION/controllers"
	

	"github.com/gorilla/mux"
)

func RegistrarModuloEducativoRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/modulo_educativo", controllers.GetModulosEducativos).Methods("GET")
	router.HandleFunc("/modulo_educativo/{id}", controllers.GetModuloEducativoByID).Methods("GET")
	router.HandleFunc("/modulo_educativo", controllers.CreateModuloEducativo).Methods("POST")
	router.HandleFunc("/modulo_educativo/{id}", controllers.UpdateModuloEducativo).Methods("PUT")
	router.HandleFunc("/modulo_educativo/{id}", controllers.DeleteModuloEducativo).Methods("DELETE")

	return router
}
