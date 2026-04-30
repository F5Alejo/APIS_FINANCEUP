package router

import (
	"AUTH/controller"

	"github.com/gorilla/mux"
)

func RegisterUsuarioRolRoutes(r *mux.Router) {
	r.HandleFunc("/auth/usuario-roles", controller.GetAllUsuarioRoles).Methods("GET")
	r.HandleFunc("/auth/usuario-roles/{id}", controller.GetUsuarioRolByID).Methods("GET")
	r.HandleFunc("/auth/usuario-roles", controller.CreateUsuarioRol).Methods("POST")
	r.HandleFunc("/auth/usuario-roles/{id}", controller.UpdateUsuarioRol).Methods("PUT")
	r.HandleFunc("/auth/usuario-roles/{id}", controller.DeleteUsuarioRol).Methods("DELETE")
}
