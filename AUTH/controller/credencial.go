package controller

import (
	"encoding/json"
	"net/http"

	"AUTH/config"
	"AUTH/models"

	"github.com/gorilla/mux"
)

// GET ALL credenciales
func GetAllCredenciales(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id_credencial, id_usuario, contrasena_hash, intentos_fallidos, requiere_cambio, activo FROM auth.credencial",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Credencial
	for rows.Next() {
		var c models.Credencial
		rows.Scan(&c.IDCredencial, &c.IDUsuario, &c.ContrasenaHash, &c.IntentosFallidos, &c.RequiereCambio, &c.Activo)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

// GET credencial by ID
func GetCredencialByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.Credencial
	err := config.DB.QueryRow(
		"SELECT id_credencial, id_usuario, contrasena_hash, intentos_fallidos, requiere_cambio, activo FROM auth.credencial WHERE id_credencial=$1", id,
	).Scan(&c.IDCredencial, &c.IDUsuario, &c.ContrasenaHash, &c.IntentosFallidos, &c.RequiereCambio, &c.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Credencial no encontrada"})
		return
	}
	respondJSON(w, 200, c)
}

// CREATE credencial
func CreateCredencial(w http.ResponseWriter, r *http.Request) {
	var c models.Credencial
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		"INSERT INTO auth.credencial (id_usuario, contrasena_hash, intentos_fallidos, requiere_cambio, activo) VALUES ($1,$2,$3,$4,$5) RETURNING id_credencial",
		c.IDUsuario, c.ContrasenaHash, c.IntentosFallidos, c.RequiereCambio, c.Activo,
	).Scan(&c.IDCredencial)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, c)
}

// UPDATE credencial
func UpdateCredencial(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c models.Credencial
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		"UPDATE auth.credencial SET contrasena_hash=$1, intentos_fallidos=$2, requiere_cambio=$3 WHERE id_credencial=$4",
		c.ContrasenaHash, c.IntentosFallidos, c.RequiereCambio, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Credencial actualizada"})
}

// DELETE credencial (lógico)
func DeleteCredencial(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"UPDATE auth.credencial SET activo=false WHERE id_credencial=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Credencial desactivada (eliminación lógica)"})
}
