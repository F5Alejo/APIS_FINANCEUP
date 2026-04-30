package config

import (
	"database/sql" //para conexiones con sql
	"fmt"
	"log" //imprimir y manejar logs
	
	_ "github.com/lib/pq" //driver postgresSQL
)

var DB *sql.DB //Instancia global de la base de datos

// connetDB establece conxión con postgreSQL

func ConnectDB() {
	// variables que necesitamos para la conexion

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "FINANCE_UP"
	schema := "auth"

psqlInfo := fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", //sslmode=disable es una variable quemada que 
	host, port, user, password, dbname, schema,

	//agarra todas las variables y las concatena todas 
)

// abrir conexion db

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	log.Fatal("Error al conectar:", err) // imprime todo en un archivo .log donde estan los errores 
}

err =db.Ping() //ping hace verificar si esta abierta la conexion a la base de datos 
if err!= nil{
	log.Fatal("No se puede conectar:", err)
}

fmt.Println("Conexion a la base de datos exitosa!")
fmt.Println("Conectalo a la db: ", dbname, " y schema: ", schema )

DB=db //ASIGANAR A CONXION GLOBAL

}
