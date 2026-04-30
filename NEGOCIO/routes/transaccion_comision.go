package routes

import (
	"NEGOCIO/controllers"
	"github.com/gorilla/mux"
)

func RegisterTransaccionRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/transacciones", controllers.GetAllTransacciones).Methods("GET")
	r.HandleFunc("/negocio/transacciones/{id}", controllers.GetTransaccionByID).Methods("GET")
	r.HandleFunc("/negocio/transacciones", controllers.CreateTransaccion).Methods("POST")
	r.HandleFunc("/negocio/transacciones/{id}", controllers.UpdateTransaccion).Methods("PUT")
	r.HandleFunc("/negocio/transacciones/{id}", controllers.DeleteTransaccion).Methods("DELETE")
}