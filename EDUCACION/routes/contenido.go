package routes

import (
	"EDUCACION/controllers"
	

	"github.com/gorilla/mux"
)

func RegistrarContenidoRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/contenido", controllers.GetContenidos).Methods("GET")
	router.HandleFunc("/contenido/{id}", controllers.GetContenidoByID).Methods("GET")
	router.HandleFunc("/contenido", controllers.CreateContenido).Methods("POST")
	router.HandleFunc("/contenido/{id}", controllers.UpdateContenido).Methods("PUT")
	router.HandleFunc("/contenido/{id}", controllers.DeleteContenido).Methods("DELETE")

	return router
}