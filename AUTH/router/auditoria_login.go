package router

import (
	"AUTH/controller"

	"github.com/gorilla/mux"
)

func RegisterAuditoriaLoginRoutes(r *mux.Router) {
	r.HandleFunc("/auth/auditoria-login", controller.GetAllAuditoriaLogin).Methods("GET")
	r.HandleFunc("/auth/auditoria-login/{id}", controller.GetAuditoriaLoginByID).Methods("GET")
	r.HandleFunc("/auth/auditoria-login", controller.CreateAuditoriaLogin).Methods("POST")
	r.HandleFunc("/auth/auditoria-login/{id}", controller.UpdateAuditoriaLogin).Methods("PUT")
	r.HandleFunc("/auth/auditoria-login/{id}", controller.DeleteAuditoriaLogin).Methods("DELETE")
}
