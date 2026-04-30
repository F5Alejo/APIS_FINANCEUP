package controller

import (
	"encoding/json"
	"net/http"

	"AUTH/config"
	"AUTH/models"

	"github.com/gorilla/mux"
)

// helper respuesta JSON (compartido en el package, definido una sola vez en credential_controller.go)

// GET ALL roles
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id_rol, nombre_rol, descripcion, activo FROM auth.rol",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Rol
	for rows.Next() {
		var rol models.Rol
		rows.Scan(&rol.IDRol, &rol.NombreRol, &rol.Descripcion, &rol.Activo)
		list = append(list, rol)
	}
	respondJSON(w, 200, list)
}

// GET rol by ID
func GetRolByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var rol models.Rol
	err := config.DB.QueryRow(
		"SELECT id_rol, nombre_rol, descripcion, activo FROM auth.rol WHERE id_rol=$1", id,
	).Scan(&rol.IDRol, &rol.NombreRol, &rol.Descripcion, &rol.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Rol no encontrado"})
		return
	}
	respondJSON(w, 200, rol)
}

// CREATE rol
func CreateRol(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol
	json.NewDecoder(r.Body).Decode(&rol)

	err := config.DB.QueryRow(
		"INSERT INTO auth.rol (nombre_rol, descripcion, activo) VALUES ($1,$2,$3) RETURNING id_rol",
		rol.NombreRol, rol.Descripcion, rol.Activo,
	).Scan(&rol.IDRol)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, rol)
}

// UPDATE rol
func UpdateRol(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var rol models.Rol
	json.NewDecoder(r.Body).Decode(&rol)

	_, err := config.DB.Exec(
		"UPDATE auth.rol SET nombre_rol=$1, descripcion=$2, activo=$3 WHERE id_rol=$4",
		rol.NombreRol, rol.Descripcion, rol.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Rol actualizado"})
}

// DELETE rol (lógico)
func DeleteRol(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"UPDATE auth.rol SET activo=false WHERE id_rol=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Rol desactivado (eliminación lógica)"})
}
