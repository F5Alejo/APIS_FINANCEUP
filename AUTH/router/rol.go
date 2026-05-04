package router

import (
	"AUTH/controller"

	"github.com/gorilla/mux"
)

func RegisterRolRoutes(r *mux.Router) {
	r.HandleFunc("/auth/roles", controller.GetAllRoles).Methods("GET")
	r.HandleFunc("/auth/roles/{id}", controller.GetRolByID).Methods("GET")
	r.HandleFunc("/auth/roles", controller.CreateRol).Methods("POST")
	r.HandleFunc("/auth/roles/{id}", controller.UpdateRol).Methods("PUT")
	r.HandleFunc("/auth/roles/{id}", controller.DeleteRol).Methods("DELETE")
}
