package routes

import (
	"soporte/controller"

	"github.com/gorilla/mux"
)

// RegisterAdjuntoRoutes registra las rutas para la tabla adjunto
func RegisterAdjuntoRoutes(r *mux.Router) {
	r.HandleFunc("/adjuntos", controller.GetAllAdjuntos).Methods("GET")
	r.HandleFunc("/adjuntos/{id}", controller.GetAdjuntoByID).Methods("GET")
	r.HandleFunc("/adjuntos", controller.CreateAdjunto).Methods("POST")
	r.HandleFunc("/adjuntos/{id}", controller.UpdateAdjunto).Methods("PUT")
	r.HandleFunc("/adjuntos/{id}", controller.DeleteAdjunto).Methods("DELETE")
}
