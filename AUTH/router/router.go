package router

import "github.com/gorilla/mux"

func RegisterAUTHRoutes(r *mux.Router) {
	RegisterAuditoriaLoginRoutes(r)
	RegisterUsuarioRoutes(r)
	RegisterUsuarioRolRoutes(r)
	RegisterCredencialRoutes(r)
	RegisterRolRoutes(r)
	RegisterTipoDocumentoRoutes(r)
}
