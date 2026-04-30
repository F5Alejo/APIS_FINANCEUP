package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establece conexión con PostgreSQL apuntando al esquema negocio
func ConnectDB() {
	host     := "localhost"
	port     := 5432
	user     := "postgres"
	password := "postgres"
	dbname   := "FINANCEUP"
	schema   := "negocio"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error al conectar:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("no se puede conectar:", err)
	}

	fmt.Println("conexion a base de datos exitosa")
	fmt.Println("conectado a la db:", dbname, "y esquema:", schema)

	DB = db
}