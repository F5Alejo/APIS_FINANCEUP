package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // driver PostgreSQL
)

var DB *sql.DB // instancia global de la base de datos

// ConnectDB establece conexión con PostgreSQL apuntando al esquema soporte
func ConnectDB() {
	host     := "localhost"
	port     := 5432
	user     := "postgres"
	password := "POSTGRES"
	dbname   := "financeup"
	schema   := "soporte"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al abrir conexión a la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("Conexión exitosa a la base de datos!")
	fmt.Println("Conectado a la db:", dbname, "y esquema:", schema)
	DB = db
}
