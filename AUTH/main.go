package main

import (
	"log"
	"net/http"

	"AUTH/config"
	"AUTH/router"

	"github.com/gorilla/mux"
)

// middleware CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	router.RegisterUsuarioRoutes(r)
	router.RegisterRolRoutes(r)
	router.RegisterCredencialRoutes(r)
	router.RegisterTipoDocumentoRoutes(r)
	router.RegisterUsuarioRolRoutes(r)
	router.RegisterAuditoriaLoginRoutes(r)

	log.Println("Servidor corriendo en el puerto 8084")
	http.ListenAndServe(":8084", enableCORS(r))
}