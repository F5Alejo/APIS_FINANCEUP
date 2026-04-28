package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllContactosAsesor(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_contacto, id_asesor, whatsapp, email, telefono,
	           disponible_desde, disponible_hasta, dias_disponibles, activo
	           FROM negocio.contacto_asesor WHERE 1=1`

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.ContactoAsesor
	for rows.Next() {
		var c models.ContactoAsesor
		rows.Scan(&c.ID_Contacto, &c.ID_Asesor, &c.Whatsapp, &c.Email,
			&c.Telefono, &c.DisponibleDesde, &c.DisponibleHasta,
			&c.DiasDisponibles, &c.Activo)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

func GetContactoAsesorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.ContactoAsesor

	err := config.DB.QueryRow(
		`SELECT id_contacto, id_asesor, whatsapp, email, telefono,
		  disponible_desde, disponible_hasta, dias_disponibles, activo
		  FROM negocio.contacto_asesor WHERE id_contacto = $1`, id,
	).Scan(&c.ID_Contacto, &c.ID_Asesor, &c.Whatsapp, &c.Email,
		&c.Telefono, &c.DisponibleDesde, &c.DisponibleHasta,
		&c.DiasDisponibles, &c.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, c)
}

func CreateContactoAsesor(w http.ResponseWriter, r *http.Request) {
	var c models.ContactoAsesor
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.contacto_asesor
		 (id_asesor, whatsapp, email, telefono,
		  disponible_desde, disponible_hasta, dias_disponibles, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id_contacto`,
		c.ID_Asesor, c.Whatsapp, c.Email, c.Telefono,
		c.DisponibleDesde, c.DisponibleHasta, c.DiasDisponibles, c.Activo,
	).Scan(&c.ID_Contacto)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, c)
}

func UpdateContactoAsesor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.ContactoAsesor
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		`UPDATE negocio.contacto_asesor
		 SET id_asesor=$1, whatsapp=$2, email=$3, telefono=$4,
		     disponible_desde=$5, disponible_hasta=$6, dias_disponibles=$7, activo=$8
		 WHERE id_contacto=$9`,
		c.ID_Asesor, c.Whatsapp, c.Email, c.Telefono,
		c.DisponibleDesde, c.DisponibleHasta, c.DiasDisponibles, c.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteContactoAsesor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec("DELETE FROM negocio.contacto_asesor WHERE id_contacto=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}