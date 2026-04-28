package main

import (
	"log"
	"net/http"

	"NEGOCIO/config"
	"NEGOCIO/routes"

	"github.com/gorilla/mux"
)

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
	config.ConnectDB()

	r := mux.NewRouter()

	// negocio
	routes.RegisterBancoRoutes(r)
	routes.RegisterProductoCrediticioRoutes(r)
	routes.RegisterAsesorRoutes(r)
	routes.RegisterContactoAsesorRoutes(r)
	routes.RegisterLeadRoutes(r)
	routes.RegisterConversacionRoutes(r)
	routes.RegisterCreditoDesembolsadoRoutes(r)
	routes.RegisterTransaccionRoutes(r)

	log.Println("servidor corriendo en el puerto 8085")
	log.Fatal(http.ListenAndServe(":8085", enableCORS(r)))
}