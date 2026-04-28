package controller

import (
	"api_soporte/config"
	"api_soporte/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllPqrs retorna todas las PQR
func GetAllPqrs(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_pqr, id_usuario, descripcion, id_estado, activo, fecha_creacion, fecha_modificacion
	          FROM soporte.pqr WHERE 1=1`

	idUsuario := r.URL.Query().Get("id_usuario")
	idEstado  := r.URL.Query().Get("id_estado")

	if idUsuario != "" {
		query += " AND id_usuario=" + idUsuario
	}
	if idEstado != "" {
		query += " AND id_estado=" + idEstado
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Pqr
	for rows.Next() {
		var p models.Pqr
		rows.Scan(&p.ID, &p.IDUsuario, &p.Descripcion, &p.IDEstado, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)
		list = append(list, p)
	}
	respondJSON(w, 200, list)
}

// GetPqrByID retorna una PQR por ID
func GetPqrByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Pqr

	err := config.DB.QueryRow(
		`SELECT id_pqr, id_usuario, descripcion, id_estado, activo, fecha_creacion, fecha_modificacion
		 FROM soporte.pqr WHERE id_pqr=$1`, id,
	).Scan(&p.ID, &p.IDUsuario, &p.Descripcion, &p.IDEstado, &p.Activo, &p.FechaCreacion, &p.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "pqr no encontrada"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, p)
}

// CreatePqr crea una nueva PQR
func CreatePqr(w http.ResponseWriter, r *http.Request) {
	var p models.Pqr
	json.NewDecoder(r.Body).Decode(&p)

	err := config.DB.QueryRow(
		`INSERT INTO soporte.pqr (id_usuario, descripcion, id_estado, activo)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id_pqr, fecha_creacion, fecha_modificacion`,
		p.IDUsuario, p.Descripcion, p.IDEstado, p.Activo,
	).Scan(&p.ID, &p.FechaCreacion, &p.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, p)
}

// UpdatePqr actualiza una PQR
func UpdatePqr(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Pqr
	json.NewDecoder(r.Body).Decode(&p)

	_, err := config.DB.Exec(
		`UPDATE soporte.pqr SET id_usuario=$1, descripcion=$2, id_estado=$3, activo=$4 WHERE id_pqr=$5`,
		p.IDUsuario, p.Descripcion, p.IDEstado, p.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "pqr actualizada"})
}

// DeletePqr elimina una PQR
func DeletePqr(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`DELETE FROM soporte.pqr WHERE id_pqr=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "pqr eliminada"})
}
