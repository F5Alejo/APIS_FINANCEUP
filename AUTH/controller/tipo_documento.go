package controller

import (
	"encoding/json"
	"net/http"

	"AUTH/config"
	"AUTH/models"

	"github.com/gorilla/mux"
)

// GET ALL tipo_documento
func GetAllTiposDocumento(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id_tipo_documento, nombre, codigo, activo FROM auth.tipo_documento",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.TipoDocumento
	for rows.Next() {
		var td models.TipoDocumento

		err := rows.Scan(&td.TipoDocumento, &td.Nombre, &td.Codigo, &td.Activo)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, td)
	}
	respondJSON(w, 200, list)
}

// GET tipo_documento by ID
func GetTipoDocumentoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var td models.TipoDocumento
	err := config.DB.QueryRow(
		"SELECT id_tipo_documento, nombre, codigo, activo FROM auth.tipo_documento WHERE id_tipo_documento=$1", id,
	).Scan(&td.TipoDocumento, &td.Nombre, &td.Codigo, &td.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Tipo de documento no encontrado"})
		return
	}
	respondJSON(w, 200, td)
}

// CREATE tipo_documento
func CreateTipoDocumento(w http.ResponseWriter, r *http.Request) {
	var td models.TipoDocumento
	json.NewDecoder(r.Body).Decode(&td)

	err := config.DB.QueryRow(
		"INSERT INTO auth.tipo_documento (nombre, codigo, activo) VALUES ($1,$2,$3) RETURNING id_tipo_documento",
		td.Nombre, td.Codigo, td.Activo,
	).Scan(&td.TipoDocumento)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, td)
}

// UPDATE tipo_documento
func UpdateTipoDocumento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var td models.TipoDocumento
	json.NewDecoder(r.Body).Decode(&td)

	_, err := config.DB.Exec(
		"UPDATE auth.tipo_documento SET nombre=$1, codigo=$2, activo=$3 WHERE id_tipo_documento=$4",
		td.Nombre, td.Codigo, td.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Tipo de documento actualizado"})
}

// DELETE tipo_documento (lógico)
func DeleteTipoDocumento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"UPDATE auth.tipo_documento SET activo=false WHERE id_tipo_documento=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Tipo de documento desactivado (eliminación lógica)"})
}
