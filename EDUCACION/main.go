package main

import (
	"log"
	"net/http"

	"EDUCACION/config"
	"EDUCACION/routes"

	"github.com/gorilla/mux"
)




func main() {
	config.ConnectDB()

	r:=mux.NewRouter()
	r.PathPrefix("/").Handler(routes.RegistrarModuloEducativoRoutes())
	r.PathPrefix("/").Handler(routes.RegistrarContenidoRoutes())


	log.Println("Servidor EDUCACION escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

