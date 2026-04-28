package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllTransacciones(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_transaccion, id_credito, id_banco, monto_comision,
	           porcentaje_aplicado, fecha_transaccion, estado, referencia_pago, activo
	           FROM negocio.transaccion_comision WHERE 1=1`

	estado := r.URL.Query().Get("estado")
	if estado != "" {
		query += " AND estado ILIKE '%" + estado + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.TransaccionComision
	for rows.Next() {
		var t models.TransaccionComision
		rows.Scan(&t.ID_Transaccion, &t.ID_Credito, &t.ID_Banco,
			&t.MontoComision, &t.PorcentajeAplicado, &t.FechaTransaccion,
			&t.Estado, &t.ReferenciaPago, &t.Activo)
		list = append(list, t)
	}
	respondJSON(w, 200, list)
}

func GetTransaccionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var t models.TransaccionComision

	err := config.DB.QueryRow(
		`SELECT id_transaccion, id_credito, id_banco, monto_comision,
		  porcentaje_aplicado, fecha_transaccion, estado, referencia_pago, activo
		  FROM negocio.transaccion_comision WHERE id_transaccion = $1`, id,
	).Scan(&t.ID_Transaccion, &t.ID_Credito, &t.ID_Banco,
		&t.MontoComision, &t.PorcentajeAplicado, &t.FechaTransaccion,
		&t.Estado, &t.ReferenciaPago, &t.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, t)
}

func CreateTransaccion(w http.ResponseWriter, r *http.Request) {
	var t models.TransaccionComision
	json.NewDecoder(r.Body).Decode(&t)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.transaccion_comision
		 (id_credito, id_banco, monto_comision, porcentaje_aplicado,
		  estado, referencia_pago, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_transaccion`,
		t.ID_Credito, t.ID_Banco, t.MontoComision, t.PorcentajeAplicado,
		t.Estado, t.ReferenciaPago, t.Activo,
	).Scan(&t.ID_Transaccion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, t)
}

func UpdateTransaccion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var t models.TransaccionComision
	json.NewDecoder(r.Body).Decode(&t)

	_, err := config.DB.Exec(
		`UPDATE negocio.transaccion_comision
		 SET id_credito=$1, id_banco=$2, monto_comision=$3,
		     porcentaje_aplicado=$4, estado=$5, referencia_pago=$6, activo=$7
		 WHERE id_transaccion=$8`,
		t.ID_Credito, t.ID_Banco, t.MontoComision,
		t.PorcentajeAplicado, t.Estado, t.ReferenciaPago, t.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteTransaccion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec("DELETE FROM negocio.transaccion_comision WHERE id_transaccion=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}