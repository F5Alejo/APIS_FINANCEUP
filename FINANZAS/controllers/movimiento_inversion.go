package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/models"

	"github.com/lib/pq"
)

func ObtenerMovimientosInversion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion FROM movimiento_inversion ORDER BY id_movimiento_dinero`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar movimiento_inversion")
		return
	}
	defer rows.Close()
	items := []models.MovimientoInversion{}
	for rows.Next() {
		var item models.MovimientoInversion
		if err := rows.Scan(&item.IdMovimientoDinero, &item.Nombre, &item.Monto, &item.EsIngreso, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer movimiento_inversion")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerMovimientoInversionPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.MovimientoInversion
	err = config.DB.QueryRow(`SELECT id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion FROM movimiento_inversion WHERE id_movimiento_dinero = $1`, id).
		Scan(&item.IdMovimientoDinero, &item.Nombre, &item.Monto, &item.EsIngreso, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "movimiento_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar movimiento_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearMovimientoInversion(w http.ResponseWriter, r *http.Request) {
	var item models.MovimientoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err := config.DB.QueryRow(`INSERT INTO movimiento_inversion (nombre, monto, es_ingreso, activo) VALUES ($1, $2, $3, $4) RETURNING id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion`,
		item.Nombre, item.Monto, item.EsIngreso, item.Activo).
		Scan(&item.IdMovimientoDinero, &item.Nombre, &item.Monto, &item.EsIngreso, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear movimiento_inversion")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarMovimientoInversion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.MovimientoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err = config.DB.QueryRow(`UPDATE movimiento_inversion SET nombre = $1, monto = $2, es_ingreso = $3, activo = $4 WHERE id_movimiento_dinero = $5 RETURNING id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion`,
		item.Nombre, item.Monto, item.EsIngreso, item.Activo, id).
		Scan(&item.IdMovimientoDinero, &item.Nombre, &item.Monto, &item.EsIngreso, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "movimiento_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar movimiento_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarMovimientoInversion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	result, err := config.DB.Exec(`DELETE FROM movimiento_inversion WHERE id_movimiento_dinero = $1`, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar movimiento_inversion porque tiene registros relacionados")
		} else {
			writeError(w, http.StatusInternalServerError, "error al eliminar movimiento_inversion")
		}
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, "movimiento_inversion no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "movimiento_inversion eliminado correctamente"})
}
