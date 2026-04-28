package router

import (
	"api_go_CRUD/controller"

	"github.com/gorilla/mux"
)

func RegisterCredencialRoutes(r *mux.Router) {
	r.HandleFunc("/auth/credenciales", controller.GetAllCredenciales).Methods("GET")
	r.HandleFunc("/auth/credenciales/{id}", controller.GetCredencialByID).Methods("GET")
	r.HandleFunc("/auth/credenciales", controller.CreateCredencial).Methods("POST")
	r.HandleFunc("/auth/credenciales/{id}", controller.UpdateCredencial).Methods("PUT")
	r.HandleFunc("/auth/credenciales/{id}", controller.DeleteCredencial).Methods("DELETE")
}
