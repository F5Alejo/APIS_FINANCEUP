package controller

import (
	"api_soporte/config"
	"api_soporte/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllEstadosPqr retorna todos los estados de PQR
func GetAllEstadosPqr(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_estado, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
	          FROM soporte.estado_pqr WHERE 1=1`

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

	var list []models.EstadoPqr
	for rows.Next() {
		var e models.EstadoPqr
		rows.Scan(&e.ID, &e.Nombre, &e.Descripcion, &e.Activo, &e.FechaCreacion, &e.FechaModificacion)
		list = append(list, e)
	}
	respondJSON(w, 200, list)
}

// GetEstadoPqrByID retorna un estado de PQR por ID
func GetEstadoPqrByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var e models.EstadoPqr

	err := config.DB.QueryRow(
		`SELECT id_estado, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
		 FROM soporte.estado_pqr WHERE id_estado=$1`, id,
	).Scan(&e.ID, &e.Nombre, &e.Descripcion, &e.Activo, &e.FechaCreacion, &e.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "estado_pqr no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, e)
}

// CreateEstadoPqr crea un nuevo estado de PQR
func CreateEstadoPqr(w http.ResponseWriter, r *http.Request) {
	var e models.EstadoPqr
	json.NewDecoder(r.Body).Decode(&e)

	err := config.DB.QueryRow(
		`INSERT INTO soporte.estado_pqr (nombre, descripcion, activo)
		 VALUES ($1, $2, $3)
		 RETURNING id_estado, fecha_creacion, fecha_modificacion`,
		e.Nombre, e.Descripcion, e.Activo,
	).Scan(&e.ID, &e.FechaCreacion, &e.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, e)
}

// UpdateEstadoPqr actualiza un estado de PQR
func UpdateEstadoPqr(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var e models.EstadoPqr
	json.NewDecoder(r.Body).Decode(&e)

	_, err := config.DB.Exec(
		`UPDATE soporte.estado_pqr SET nombre=$1, descripcion=$2, activo=$3 WHERE id_estado=$4`,
		e.Nombre, e.Descripcion, e.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "estado_pqr actualizado"})
}

// DeleteEstadoPqr elimina un estado de PQR
func DeleteEstadoPqr(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`DELETE FROM soporte.estado_pqr WHERE id_estado=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "estado_pqr eliminado"})
}
