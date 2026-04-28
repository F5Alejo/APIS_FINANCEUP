package controller

import (
	"soporte/config"
	"soporte/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllRegistrosActividad retorna todos los registros de actividad
func GetAllRegistrosActividad(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_actividad, id_usuario, tipo_actividad, descripcion, entidad_afectada,
	          fecha_actividad, fecha_creacion
	          FROM soporte.registro_actividad WHERE 1=1`

	idUsuario       := r.URL.Query().Get("id_usuario")
	tipoActividad   := r.URL.Query().Get("tipo_actividad")
	entidadAfectada := r.URL.Query().Get("entidad_afectada")

	if idUsuario != "" {
		query += " AND id_usuario=" + idUsuario
	}
	if tipoActividad != "" {
		query += " AND tipo_actividad ILIKE '%" + tipoActividad + "%'"
	}
	if entidadAfectada != "" {
		query += " AND entidad_afectada ILIKE '%" + entidadAfectada + "%'"
	}

	query += " ORDER BY fecha_actividad DESC"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.RegistroActividad
	for rows.Next() {
		var ra models.RegistroActividad
		rows.Scan(&ra.ID, &ra.IDUsuario, &ra.TipoActividad, &ra.Descripcion,
			&ra.EntidadAfectada, &ra.FechaActividad, &ra.FechaCreacion)
		list = append(list, ra)
	}
	respondJSON(w, 200, list)
}

// GetRegistroActividadByID retorna un registro de actividad por ID
func GetRegistroActividadByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var ra models.RegistroActividad

	err := config.DB.QueryRow(
		`SELECT id_actividad, id_usuario, tipo_actividad, descripcion, entidad_afectada,
		 fecha_actividad, fecha_creacion
		 FROM soporte.registro_actividad WHERE id_actividad=$1`, id,
	).Scan(&ra.ID, &ra.IDUsuario, &ra.TipoActividad, &ra.Descripcion,
		&ra.EntidadAfectada, &ra.FechaActividad, &ra.FechaCreacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "registro de actividad no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, ra)
}

// CreateRegistroActividad crea un nuevo registro de actividad
// Nota: registro_actividad es inmutable por diseño (no tiene UPDATE ni DELETE)
func CreateRegistroActividad(w http.ResponseWriter, r *http.Request) {
	var ra models.RegistroActividad
	json.NewDecoder(r.Body).Decode(&ra)

	err := config.DB.QueryRow(
		`INSERT INTO soporte.registro_actividad (id_usuario, tipo_actividad, descripcion, entidad_afectada)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id_actividad, fecha_actividad, fecha_creacion`,
		ra.IDUsuario, ra.TipoActividad, ra.Descripcion, ra.EntidadAfectada,
	).Scan(&ra.ID, &ra.FechaActividad, &ra.FechaCreacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, ra)
}
