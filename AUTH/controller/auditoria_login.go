package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"AUTH/config"
	"AUTH/models"

	"github.com/gorilla/mux"
)

// GET ALL auditoria_login con filtros opcionales
func GetAllAuditoriaLogin(w http.ResponseWriter, r *http.Request) {
	baseQuery := "SELECT id_auditoria, id_usuario, estado_evento, navegador, fecha_evento FROM auth.auditoria_login WHERE 1=1"
	var args []interface{}
	argCount := 1

	idUsuario := r.URL.Query().Get("id_usuario")
	estadoEvento := r.URL.Query().Get("estado_evento")

	if idUsuario != "" {
		baseQuery += " AND id_usuario=$" + strconv.Itoa(argCount)
		args = append(args, idUsuario)
		argCount++
	}
	if estadoEvento != "" {
		baseQuery += " AND estado_evento=$" + strconv.Itoa(argCount)
		args = append(args, estadoEvento)
		argCount++
	}

	rows, err := config.DB.Query(baseQuery, args...)
	if err != nil {
		log.Printf("Error al ejecutar la consulta de auditoría: %v", err) // Agregado para depuración
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.AuditoriaLogin
	for rows.Next() {
		var a models.AuditoriaLogin
		rows.Scan(&a.IDAuditoria, &a.IDUsuario, &a.EstadoEvento, &a.Navegador, &a.FechaEvento)
		list = append(list, a)
	}
	respondJSON(w, 200, list)
}

// GET auditoria_login by ID
func GetAuditoriaLoginByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var a models.AuditoriaLogin
	err := config.DB.QueryRow(
		"SELECT id_auditoria, id_usuario, estado_evento, navegador, fecha_evento FROM auth.auditoria_login WHERE id_auditoria=$1", id,
	).Scan(&a.IDAuditoria, &a.IDUsuario, &a.EstadoEvento, &a.Navegador, &a.FechaEvento)

	if err != nil {
		if err == sql.ErrNoRows {
			respondJSON(w, 404, map[string]string{"error": "Registro de auditoría no encontrado"})
		} else {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
		}
		return
	}
	respondJSON(w, 200, a)
}

// CREATE auditoria_login
func CreateAuditoriaLogin(w http.ResponseWriter, r *http.Request) {
	var a models.AuditoriaLogin
	json.NewDecoder(r.Body).Decode(&a)

	err := config.DB.QueryRow(
		"INSERT INTO auth.auditoria_login (id_usuario, estado_evento, navegador) VALUES ($1,$2,$3) RETURNING id_auditoria",
		a.IDUsuario, a.EstadoEvento, a.Navegador,
	).Scan(&a.IDAuditoria)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, a)
}

// UPDATE auditoria_login
func UpdateAuditoriaLogin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var a models.AuditoriaLogin
	json.NewDecoder(r.Body).Decode(&a)

	_, err := config.DB.Exec(
		"UPDATE auth.auditoria_login SET estado_evento=$1, navegador=$2 WHERE id_auditoria=$3",
		a.EstadoEvento, a.Navegador, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Registro de auditoría actualizado"})
}

// DELETE auditoria_login (físico)
func DeleteAuditoriaLogin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec(
		"DELETE FROM auth.auditoria_login WHERE id_auditoria=$1", id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Registro de auditoría eliminado"})
}
