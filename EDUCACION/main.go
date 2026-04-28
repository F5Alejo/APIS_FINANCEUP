package main

import (
	"log"
	"net/http"

	"EDUCACION/config"
	"EDUCACION/routes"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRoutes()

	log.Println("Servidor EDUCACION escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

