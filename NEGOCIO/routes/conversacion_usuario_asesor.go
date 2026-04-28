package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterConversacionRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/conversaciones", controllers.GetAllConversaciones).Methods("GET")
	r.HandleFunc("/negocio/conversaciones/{id}", controllers.GetConversacionByID).Methods("GET")
	r.HandleFunc("/negocio/conversaciones", controllers.CreateConversacion).Methods("POST")
	r.HandleFunc("/negocio/conversaciones/{id}", controllers.UpdateConversacion).Methods("PUT")
	r.HandleFunc("/negocio/conversaciones/{id}", controllers.DeleteConversacion).Methods("DELETE")
}