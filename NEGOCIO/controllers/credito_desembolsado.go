package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCreditosDesembolsados(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_credito, id_lead, id_usuario, id_producto, id_banco,
	           numero_credito, monto_aprobado, tasa_interes_final, plazo_meses,
	           fecha_aprobacion, fecha_desembolso, estado_credito, saldo_actual, activo
	           FROM negocio.credito_desembolsado WHERE 1=1`

	estado := r.URL.Query().Get("estado_credito")
	if estado != "" {
		query += " AND estado_credito ILIKE '%" + estado + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.CreditoDesembolsado
	for rows.Next() {
		var c models.CreditoDesembolsado
		rows.Scan(&c.ID_Credito, &c.ID_Lead, &c.ID_Usuario, &c.ID_Producto, &c.ID_Banco,
			&c.NumeroCredito, &c.MontoAprobado, &c.TasaInterestFinal, &c.PlazoMeses,
			&c.FechaAprobacion, &c.FechaDesembolso, &c.EstadoCredito, &c.SaldoActual, &c.Activo)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

func GetCreditoDesembolsadoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.CreditoDesembolsado

	err := config.DB.QueryRow(
		`SELECT id_credito, id_lead, id_usuario, id_producto, id_banco,
		  numero_credito, monto_aprobado, tasa_interes_final, plazo_meses,
		  fecha_aprobacion, fecha_desembolso, estado_credito, saldo_actual, activo
		  FROM negocio.credito_desembolsado WHERE id_credito = $1`, id,
	).Scan(&c.ID_Credito, &c.ID_Lead, &c.ID_Usuario, &c.ID_Producto, &c.ID_Banco,
		&c.NumeroCredito, &c.MontoAprobado, &c.TasaInterestFinal, &c.PlazoMeses,
		&c.FechaAprobacion, &c.FechaDesembolso, &c.EstadoCredito, &c.SaldoActual, &c.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, c)
}

func CreateCreditoDesembolsado(w http.ResponseWriter, r *http.Request) {
	var c models.CreditoDesembolsado
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.credito_desembolsado
		 (id_lead, id_usuario, id_producto, id_banco, numero_credito,
		  monto_aprobado, tasa_interes_final, plazo_meses,
		  fecha_aprobacion, fecha_desembolso, estado_credito, saldo_actual, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id_credito`,
		c.ID_Lead, c.ID_Usuario, c.ID_Producto, c.ID_Banco, c.NumeroCredito,
		c.MontoAprobado, c.TasaInterestFinal, c.PlazoMeses,
		c.FechaAprobacion, c.FechaDesembolso, c.EstadoCredito, c.SaldoActual, c.Activo,
	).Scan(&c.ID_Credito)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, c)
}

func UpdateCreditoDesembolsado(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.CreditoDesembolsado
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		`UPDATE negocio.credito_desembolsado
		 SET id_lead=$1, id_usuario=$2, id_producto=$3, id_banco=$4,
		     numero_credito=$5, monto_aprobado=$6, tasa_interes_final=$7,
		     plazo_meses=$8, fecha_aprobacion=$9, fecha_desembolso=$10,
		     estado_credito=$11, saldo_actual=$12, activo=$13
		 WHERE id_credito=$14`,
		c.ID_Lead, c.ID_Usuario, c.ID_Producto, c.ID_Banco,
		c.NumeroCredito, c.MontoAprobado, c.TasaInterestFinal,
		c.PlazoMeses, c.FechaAprobacion, c.FechaDesembolso,
		c.EstadoCredito, c.SaldoActual, c.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteCreditoDesembolsado(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec("DELETE FROM negocio.credito_desembolsado WHERE id_credito=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}