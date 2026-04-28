package controllers

import (
	"API_GO_CRUD/config"
	"API_GO_CRUD/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAsesores(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_asesor, id_banco, nombre, apellido, email, telefono, especialidad, activo
	          FROM negocio.asesor_bancario WHERE 1=1`

	nombre := r.URL.Query().Get("nombre")
	if nombre != "" {
		query += " AND nombre ILIKE '%" + nombre + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.AsesorBancario
	for rows.Next() {
		var a models.AsesorBancario
		rows.Scan(&a.ID, &a.IDBanco, &a.Nombre, &a.Apellido, &a.Email, &a.Telefono, &a.Especialidad, &a.Activo)
		list = append(list, a)
	}
	respondJSON(w, 200, list)
}

func GetAsesorByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.AsesorBancario

	err := config.DB.QueryRow(
		`SELECT id_asesor, id_banco, nombre, apellido, email, telefono, especialidad, activo
		 FROM negocio.asesor_bancario WHERE id_asesor = $1`, id,
	).Scan(&a.ID, &a.IDBanco, &a.Nombre, &a.Apellido, &a.Email, &a.Telefono, &a.Especialidad, &a.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "no encontrado"})
		return
	}
	respondJSON(w, 200, a)
}

func CreateAsesor(w http.ResponseWriter, r *http.Request) {
	var a models.AsesorBancario
	json.NewDecoder(r.Body).Decode(&a)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.asesor_bancario (id_banco, nombre, apellido, email, telefono, especialidad, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_asesor`,
		a.IDBanco, a.Nombre, a.Apellido, a.Email, a.Telefono, a.Especialidad, a.Activo,
	).Scan(&a.ID)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, a)
}

func UpdateAsesor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.AsesorBancario
	json.NewDecoder(r.Body).Decode(&a)

	_, err := config.DB.Exec(
		`UPDATE negocio.asesor_bancario
		 SET id_banco=$1, nombre=$2, apellido=$3, email=$4, telefono=$5, especialidad=$6, activo=$7
		 WHERE id_asesor=$8`,
		a.IDBanco, a.Nombre, a.Apellido, a.Email, a.Telefono, a.Especialidad, a.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteAsesor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM negocio.asesor_bancario WHERE id_asesor=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}