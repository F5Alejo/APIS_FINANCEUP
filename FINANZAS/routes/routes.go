package routes

import "github.com/gorilla/mux"

func RegistrarRutas(router *mux.Router) {
	RegistrarRutasCategoria(router)
	RegistrarRutasMovimientoIngresoEgreso(router)
	RegistrarRutasTipoIngreso(router)
	RegistrarRutasFinanzas(router)
	RegistrarRutasTipoInversion(router)
	RegistrarRutasNivelRiesgo(router)
	RegistrarRutasMovimientoInversion(router)
	RegistrarRutasTipoIngresoInversion(router)
	RegistrarRutasInversion(router)
	RegistrarRutasEditarMeta(router)
	RegistrarRutasMovimientoMeta(router)
	RegistrarRutasTipoIngresoMeta(router)
	RegistrarRutasMeta(router)
}
