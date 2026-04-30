package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterLeadRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/leads", controllers.GetAllLeads).Methods("GET")
	r.HandleFunc("/negocio/leads/{id}", controllers.GetLeadByID).Methods("GET")
	r.HandleFunc("/negocio/leads", controllers.CreateLead).Methods("POST")
	r.HandleFunc("/negocio/leads/{id}", controllers.UpdateLead).Methods("PUT")
	r.HandleFunc("/negocio/leads/{id}", controllers.DeleteLead).Methods("DELETE")
}