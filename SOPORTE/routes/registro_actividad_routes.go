package routes

import (
	"api_soporte/controller"

	"github.com/gorilla/mux"
)

// RegisterRegistroActividadRoutes registra las rutas para registro_actividad
// Nota: solo GET y POST — los registros de auditoría son inmutables por diseño del esquema
func RegisterRegistroActividadRoutes(r *mux.Router) {
	r.HandleFunc("/registros-actividad", controller.GetAllRegistrosActividad).Methods("GET")
	r.HandleFunc("/registros-actividad/{id}", controller.GetRegistroActividadByID).Methods("GET")
	r.HandleFunc("/registros-actividad", controller.CreateRegistroActividad).Methods("POST")
}
