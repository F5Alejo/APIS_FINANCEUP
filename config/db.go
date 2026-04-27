package config

import (
	"database/sql" // para conecciones con sql
	"fmt"
	"log" // imprimir y manejar logs

	//variables de entorno
	_ "github.com/lib/pq" //driver postgresSQL
)

var DB *sql.DB //insncia global de la base de datos

//connectDB establece conexIon con postgresSQL

func ConnectDB() {
	//varibles para la conexion
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "transversal10"
	dbname := "go_db"
	schema := "go_api"


	// CADENA DE CONEXION
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)
	// abrir conexion db

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error al conectar:", err)

	}

	err = db.Ping()
	if err != nil {
		log.Fatal("no se puede conectar:", err)

	}

	fmt.Println("conexion a base de datos exitosa")
	fmt.Println("Conectado a la db: ", dbname,"y esquema: ",schema)
	DB = db // ASIGNAR A CONEXION GLOBAL

}

//cadena de conexion
