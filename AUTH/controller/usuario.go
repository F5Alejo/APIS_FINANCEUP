package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"api_go_CRUD/config"
	"api_go_CRUD/models"

	"github.com/gorilla/mux"
)

// GET ALL usuarios con filtros opcionales por nombre, email y estado
func GetAllUsuarios(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_usuario, nombre, apellido, email, cedula, ciudad, estado, id_tipo_documento, activo FROM auth.usuario WHERE 1=1"

	nombre := r.URL.Query().Get("nombre")
	email := r.URL.Query().Get("email")
	estado := r.URL.Query().Get("estado")

	if nombre != "" {
		query += " AND nombre ILIKE '%" + nombre + "%'"
	}
	if email != "" {
		query += " AND email ILIKE '%" + email + "%'"
	}
	if estado != "" {
		query += " AND estado = '" + estado + "'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var u models.Usuario
		rows.Scan(&u.IDUsuario, &u.Nombre, &u.Apellido, &u.Email, &u.Cedula, &u.Ciudad, &u.Estado, &u.IDTipoDocumento, &u.Activo)
		usuarios = append(usuarios, u)
	}
	respondJSON(w, 200, usuarios)
}

// GET usuario by ID
func GetUsuarioByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var u models.Usuario
	err := config.DB.QueryRow(
		"SELECT id_usuario, nombre, apellido, email, cedula, ciudad, estado, id_tipo_documento, activo FROM auth.usuario WHERE id_usuario=$1", id,
	).Scan(&u.IDUsuario, &u.Nombre, &u.Apellido, &u.Email, &u.Cedula, &u.Ciudad, &u.Estado, &u.IDTipoDocumento, &u.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, u)
}

// CREATE usuario
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	err := config.DB.QueryRow(
		"INSERT INTO auth.usuario (nombre, apellido, email, cedula, ciudad, estado, id_tipo_documento, activo) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id_usuario",
		u.Nombre, u.Apellido, u.Email, u.Cedula, u.Ciudad, u.Estado, u.IDTipoDocumento, u.Activo,
	).Scan(&u.IDUsuario)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, u)
}

// UPDATE usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var u models.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	_, err := config.DB.Exec(
		"UPDATE auth.usuario SET nombre=$1, apellido=$2, email=$3, cedula=$4, ciudad=$5, estado=$6, id_tipo_documento=$7, activo=$8 WHERE id_usuario=$9",
		u.Nombre, u.Apellido, u.Email, u.Cedula, u.Ciudad, u.Estado, u.IDTipoDocumento, u.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, u)
}

// DELETE usuario (lógico)
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"UPDATE auth.usuario SET activo=false WHERE id_usuario=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Usuario desactivado (eliminación lógica)"})
}
