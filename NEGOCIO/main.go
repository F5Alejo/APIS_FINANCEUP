package main

import (
	"log"
	"net/http"

	"FINANCEUP/config"
	"FINANCEUP/routes"

	"github.com/gorilla/mux"
)

// enableCORS middleware para permitir peticiones de cualquier origen
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

	// ── Registro de rutas del esquema NEGOCIO ──
	routes.RegisterBancoRoutes(r)
	routes.RegisterProductoCrediticioRoutes(r)
	routes.RegisterAsesorBancarioRoutes(r)
	routes.RegisterContactoAsesorRoutes(r)
	routes.RegisterLeadRoutes(r)
	routes.RegisterConversacionRoutes(r)
	routes.RegisterCreditoDesembolsadoRoutes(r)
	routes.RegisterTransaccionComisionRoutes(r)

	log.Println("servidor FINANCEUP - esquema NEGOCIO corriendo en el puerto 8085")
	log.Fatal(http.ListenAndServe(":8085", enableCORS(r)))
}