package routes

import (
	"FINANZAS/controllers"

	"github.com/gorilla/mux"
)

func RegistrarRutasMeta(router *mux.Router) {
	router.HandleFunc("/meta", controllers.ObtenerMetas).Methods("GET")
	router.HandleFunc("/meta/{id}", controllers.ObtenerMetaPorID).Methods("GET")
	router.HandleFunc("/meta", controllers.CrearMeta).Methods("POST")
	router.HandleFunc("/meta/{id}", controllers.ActualizarMeta).Methods("PUT")
	router.HandleFunc("/meta/{id}", controllers.EliminarMeta).Methods("DELETE")
}
