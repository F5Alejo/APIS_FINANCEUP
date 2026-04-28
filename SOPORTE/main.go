package main

import (
	"log"
	"net/http"

	"api_soporte/config"
	"api_soporte/routes"

	"github.com/gorilla/mux"
)

// enableCORS middleware para permitir peticiones cross-origin
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB() // conectar DB

	r := mux.NewRouter()

	// Registro de rutas — esquema soporte
	routes.RegisterPqrRoutes(r)

	log.Println("Servidor soporte corriendo en el puerto 8080")
	http.ListenAndServe(":8080", enableCORS(r))
}
