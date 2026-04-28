package routes

import "github.com/gorilla/mux"

func RegistrarRutas(router *mux.Router) {
	RegistrarRutasCategoria(router)
}
