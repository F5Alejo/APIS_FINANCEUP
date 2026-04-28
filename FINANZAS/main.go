package main

import (
	"log"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegistrarRutas(r)

	log.Println("Servidor FINANZAS escuchando en :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
