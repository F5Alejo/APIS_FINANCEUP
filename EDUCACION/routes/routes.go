package routes

import "github.com/gorilla/mux"

func RegistrarRutas(router *mux.Router) {
	RegistrarRutasModuloEducativo(router)
	RegistrarRutasContenido(router)
	RegistrarRutasLeccion(router)
	RegistrarRutasProgresoEducativo(router)
	RegistrarRutasProgresoLeccion(router)
}
