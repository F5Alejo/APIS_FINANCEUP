package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllBancos obtiene todos los bancos con filtros opcionales
func GetAllBancos(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_banco, nombre_banco, ciudad, contacto, telefono, email, comision_porcentaje, url_logo, descripcion, sitio_web, estado, activo FROM banco WHERE 1=1"

	nombre := r.URL.Query().Get("nombre_banco")
	estado  := r.URL.Query().Get("estado")

	if nombre != "" {
		query += " AND nombre_banco ILIKE '%" + nombre + "%'"
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

	var bancos []models.Banco
	for rows.Next() {
		var b models.Banco
		rows.Scan(
			&b.IDBanco, &b.NombreBanco, &b.Ciudad, &b.Contacto,
			&b.Telefono, &b.Email, &b.ComisionPorcentaje, &b.UrlLogo,
			&b.Descripcion, &b.SitioWeb, &b.Estado, &b.Activo,
		)
		bancos = append(bancos, b)
	}

	respondJSON(w, 200, bancos)
}

// GetBancoByID obtiene un banco por su ID
func GetBancoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var b models.Banco

	err := config.DB.QueryRow(
		"SELECT id_banco, nombre_banco, ciudad, contacto, telefono, email, comision_porcentaje, url_logo, descripcion, sitio_web, estado, activo FROM banco WHERE id_banco=$1",
		id,
	).Scan(
		&b.IDBanco, &b.NombreBanco, &b.Ciudad, &b.Contacto,
		&b.Telefono, &b.Email, &b.ComisionPorcentaje, &b.UrlLogo,
		&b.Descripcion, &b.SitioWeb, &b.Estado, &b.Activo,
	)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "banco no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, b)
}

// CreateBanco crea un nuevo banco
func CreateBanco(w http.ResponseWriter, r *http.Request) {
	var b models.Banco
	json.NewDecoder(r.Body).Decode(&b)

	err := config.DB.QueryRow(
		`INSERT INTO banco
			(nombre_banco, ciudad, contacto, telefono, email, comision_porcentaje, url_logo, descripcion, sitio_web, estado, activo)
		VALUES
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id_banco`,
		b.NombreBanco, b.Ciudad, b.Contacto, b.Telefono, b.Email,
		b.ComisionPorcentaje, b.UrlLogo, b.Descripcion, b.SitioWeb, b.Estado, b.Activo,
	).Scan(&b.IDBanco)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, b)
}

// UpdateBanco actualiza un banco por ID
func UpdateBanco(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var b models.Banco
	json.NewDecoder(r.Body).Decode(&b)

	_, err := config.DB.Exec(
		`UPDATE banco SET
			nombre_banco=$1, ciudad=$2, contacto=$3, telefono=$4, email=$5,
			comision_porcentaje=$6, url_logo=$7, descripcion=$8, sitio_web=$9,
			estado=$10, activo=$11
		WHERE id_banco=$12`,
		b.NombreBanco, b.Ciudad, b.Contacto, b.Telefono, b.Email,
		b.ComisionPorcentaje, b.UrlLogo, b.Descripcion, b.SitioWeb,
		b.Estado, b.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Banco actualizado"})
}

// DeleteBanco elimina un banco por ID
func DeleteBanco(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM banco WHERE id_banco=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Banco eliminado"})
}