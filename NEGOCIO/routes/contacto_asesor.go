package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterContactoAsesorRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/contactos", controllers.GetAllContactosAsesor).Methods("GET")
	r.HandleFunc("/negocio/contactos/{id}", controllers.GetContactoAsesorByID).Methods("GET")
	r.HandleFunc("/negocio/contactos", controllers.CreateContactoAsesor).Methods("POST")
	r.HandleFunc("/negocio/contactos/{id}", controllers.UpdateContactoAsesor).Methods("PUT")
	r.HandleFunc("/negocio/contactos/{id}", controllers.DeleteContactoAsesor).Methods("DELETE")
}