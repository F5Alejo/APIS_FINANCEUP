package router

import (
	"api_go_CRUD/controller"

	"github.com/gorilla/mux"
)

func RegisterUsuarioRoutes(r *mux.Router) {
	r.HandleFunc("/auth/usuarios", controller.GetAllUsuarios).Methods("GET")
	r.HandleFunc("/auth/usuarios/{id}", controller.GetUsuarioByID).Methods("GET")
	r.HandleFunc("/auth/usuarios", controller.CreateUsuario).Methods("POST")
	r.HandleFunc("/auth/usuarios/{id}", controller.UpdateUsuario).Methods("PUT")
	r.HandleFunc("/auth/usuarios/{id}", controller.DeleteUsuario).Methods("DELETE")
}
