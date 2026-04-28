package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllLeads(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_lead, id_usuario, id_producto, id_asesor, tipo_credito,
	           monto_interes, plazo_interes, estado_lead, fecha_generacion,
	           fecha_contacto, observaciones, activo
	           FROM negocio.lead WHERE 1=1`

	estado := r.URL.Query().Get("estado_lead")
	if estado != "" {
		query += " AND estado_lead ILIKE '%" + estado + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Lead
	for rows.Next() {
		var l models.Lead
		rows.Scan(&l.ID_Lead, &l.ID_Usuario, &l.ID_Producto, &l.ID_Asesor,
			&l.TipoCredito, &l.MontoInteres, &l.PlazoInteres, &l.EstadoLead,
			&l.FechaGeneracion, &l.FechaContacto, &l.Observaciones, &l.Activo)
		list = append(list, l)
	}
	respondJSON(w, 200, list)
}

func GetLeadByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var l models.Lead

	err := config.DB.QueryRow(
		`SELECT id_lead, id_usuario, id_producto, id_asesor, tipo_credito,
		  monto_interes, plazo_interes, estado_lead, fecha_generacion,
		  fecha_contacto, observaciones, activo
		  FROM negocio.lead WHERE id_lead = $1`, id,
	).Scan(&l.ID_Lead, &l.ID_Usuario, &l.ID_Producto, &l.ID_Asesor,
		&l.TipoCredito, &l.MontoInteres, &l.PlazoInteres, &l.EstadoLead,
		&l.FechaGeneracion, &l.FechaContacto, &l.Observaciones, &l.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, l)
}

func CreateLead(w http.ResponseWriter, r *http.Request) {
	var l models.Lead
	json.NewDecoder(r.Body).Decode(&l)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.lead
		 (id_usuario, id_producto, id_asesor, tipo_credito, monto_interes,
		  plazo_interes, estado_lead, fecha_contacto, observaciones, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id_lead`,
		l.ID_Usuario, l.ID_Producto, l.ID_Asesor, l.TipoCredito, l.MontoInteres,
		l.PlazoInteres, l.EstadoLead, l.FechaContacto, l.Observaciones, l.Activo,
	).Scan(&l.ID_Lead)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, l)
}

func UpdateLead(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var l models.Lead
	json.NewDecoder(r.Body).Decode(&l)

	_, err := config.DB.Exec(
		`UPDATE negocio.lead
		 SET id_usuario=$1, id_producto=$2, id_asesor=$3, tipo_credito=$4,
		     monto_interes=$5, plazo_interes=$6, estado_lead=$7,
		     fecha_contacto=$8, observaciones=$9, activo=$10
		 WHERE id_lead=$11`,
		l.ID_Usuario, l.ID_Producto, l.ID_Asesor, l.TipoCredito,
		l.MontoInteres, l.PlazoInteres, l.EstadoLead,
		l.FechaContacto, l.Observaciones, l.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteLead(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec("DELETE FROM negocio.lead WHERE id_lead=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}