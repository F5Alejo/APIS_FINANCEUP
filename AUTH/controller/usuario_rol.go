package controller

import (
	"encoding/json"
	"net/http"

	"AUTH/config"
	"AUTH/models"

	"github.com/gorilla/mux"
)

// GET ALL usuario_rol
func GetAllUsuarioRoles(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id_usuario_rol, id_usuario, id_rol, fecha_asignacion, activo FROM auth.usuario_rol",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.UsuarioRol
	for rows.Next() {
		var ur models.UsuarioRol
		rows.Scan(&ur.IDUsuarioRol, &ur.IDUsuario, &ur.IDRol, &ur.FechaAsignacion, &ur.Activo)
		list = append(list, ur)
	}
	respondJSON(w, 200, list)
}

// GET usuario_rol by ID
func GetUsuarioRolByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var ur models.UsuarioRol
	err := config.DB.QueryRow(
		"SELECT id_usuario_rol, id_usuario, id_rol, fecha_asignacion, activo FROM auth.usuario_rol WHERE id_usuario_rol=$1", id,
	).Scan(&ur.IDUsuarioRol, &ur.IDUsuario, &ur.IDRol, &ur.FechaAsignacion, &ur.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Asignación no encontrada"})
		return
	}
	respondJSON(w, 200, ur)
}

// CREATE usuario_rol
func CreateUsuarioRol(w http.ResponseWriter, r *http.Request) {
	var ur models.UsuarioRol
	json.NewDecoder(r.Body).Decode(&ur)

	err := config.DB.QueryRow(
		"INSERT INTO auth.usuario_rol (id_usuario, id_rol, activo) VALUES ($1,$2,$3) RETURNING id_usuario_rol",
		ur.IDUsuario, ur.IDRol, ur.Activo,
	).Scan(&ur.IDUsuarioRol)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, ur)
}

// UPDATE usuario_rol
func UpdateUsuarioRol(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var ur models.UsuarioRol
	json.NewDecoder(r.Body).Decode(&ur)

	_, err := config.DB.Exec(
		"UPDATE auth.usuario_rol SET id_rol=$1, activo=$2 WHERE id_usuario_rol=$3",
		ur.IDRol, ur.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Asignación de rol actualizada"})
}

// DELETE usuario_rol (lógico)
func DeleteUsuarioRol(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"UPDATE auth.usuario_rol SET activo=false WHERE id_usuario_rol=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Asignación desactivada (eliminación lógica)"})
}
