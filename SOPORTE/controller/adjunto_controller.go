package controller

import (
	"soporte/config"
	"soporte/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// respondJSON es el helper de respuesta JSON (compartido en el paquete)
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// GetAllAdjuntos retorna todos los adjuntos
func GetAllAdjuntos(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_adjunto, id_pqr, nombre_archivo, ruta_archivo, tipo_mime, tamano_bytes,
	          fecha_carga, activo, fecha_creacion, fecha_modificacion
	          FROM soporte.adjunto WHERE 1=1`

	idPqr := r.URL.Query().Get("id_pqr")
	if idPqr != "" {
		query += " AND id_pqr=" + idPqr
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Adjunto
	for rows.Next() {
		var a models.Adjunto
		rows.Scan(&a.ID, &a.IDPqr, &a.NombreArchivo, &a.RutaArchivo, &a.TipoMime,
			&a.TamanoBytes, &a.FechaCarga, &a.Activo, &a.FechaCreacion, &a.FechaModificacion)
		list = append(list, a)
	}
	respondJSON(w, 200, list)
}

// GetAdjuntoByID retorna un adjunto por ID
func GetAdjuntoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Adjunto

	err := config.DB.QueryRow(
		`SELECT id_adjunto, id_pqr, nombre_archivo, ruta_archivo, tipo_mime, tamano_bytes,
		 fecha_carga, activo, fecha_creacion, fecha_modificacion
		 FROM soporte.adjunto WHERE id_adjunto=$1`, id,
	).Scan(&a.ID, &a.IDPqr, &a.NombreArchivo, &a.RutaArchivo, &a.TipoMime,
		&a.TamanoBytes, &a.FechaCarga, &a.Activo, &a.FechaCreacion, &a.FechaModificacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "adjunto no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, a)
}

// CreateAdjunto crea un nuevo adjunto
func CreateAdjunto(w http.ResponseWriter, r *http.Request) {
	var a models.Adjunto
	json.NewDecoder(r.Body).Decode(&a)

	err := config.DB.QueryRow(
		`INSERT INTO soporte.adjunto (id_pqr, nombre_archivo, ruta_archivo, tipo_mime, tamano_bytes, activo)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id_adjunto, fecha_carga, fecha_creacion, fecha_modificacion`,
		a.IDPqr, a.NombreArchivo, a.RutaArchivo, a.TipoMime, a.TamanoBytes, a.Activo,
	).Scan(&a.ID, &a.FechaCarga, &a.FechaCreacion, &a.FechaModificacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, a)
}

// UpdateAdjunto actualiza un adjunto
func UpdateAdjunto(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var a models.Adjunto
	json.NewDecoder(r.Body).Decode(&a)

	_, err := config.DB.Exec(
		`UPDATE soporte.adjunto SET id_pqr=$1, nombre_archivo=$2, ruta_archivo=$3,
		 tipo_mime=$4, tamano_bytes=$5, activo=$6 WHERE id_adjunto=$7`,
		a.IDPqr, a.NombreArchivo, a.RutaArchivo, a.TipoMime, a.TamanoBytes, a.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "adjunto actualizado"})
}

// DeleteAdjunto elimina un adjunto
func DeleteAdjunto(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(`DELETE FROM soporte.adjunto WHERE id_adjunto=$1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "adjunto eliminado"})
}
