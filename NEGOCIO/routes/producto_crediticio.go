package routes

import (
	"API_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterProductoCrediticioRoutes(r *mux.Router) {
	r.HandleFunc("/negocio/productos", controllers.GetAllProductosCrediticios).Methods("GET")
	r.HandleFunc("/negocio/productos/{id}", controllers.GetProductoCrediticioByID).Methods("GET")
	r.HandleFunc("/negocio/productos", controllers.CreateProductoCrediticio).Methods("POST")
	r.HandleFunc("/negocio/productos/{id}", controllers.UpdateProductoCrediticio).Methods("PUT")
	r.HandleFunc("/negocio/productos/{id}", controllers.DeleteProductoCrediticio).Methods("DELETE")
}