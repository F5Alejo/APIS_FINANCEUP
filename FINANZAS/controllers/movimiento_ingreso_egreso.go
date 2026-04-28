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

func ObtenerMovimientosIngresoEgreso(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion FROM movimiento_ingreso_egreso ORDER BY id_movimiento_dinero`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar movimiento_ingreso_egreso")
		return
	}
	defer rows.Close()

	movimientos := []models.MovimientoIngresoEgreso{}
	for rows.Next() {
		var movimiento models.MovimientoIngresoEgreso
		if err := rows.Scan(&movimiento.IdMovimientoDinero, &movimiento.Nombre, &movimiento.Monto, &movimiento.EsIngreso, &movimiento.Activo, &movimiento.FechaCreacion, &movimiento.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer movimiento_ingreso_egreso")
			return
		}
		movimientos = append(movimientos, movimiento)
	}

	writeJSON(w, http.StatusOK, movimientos)
}

func ObtenerMovimientoIngresoEgresoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var movimiento models.MovimientoIngresoEgreso
	err = config.DB.QueryRow(`SELECT id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion FROM movimiento_ingreso_egreso WHERE id_movimiento_dinero = $1`, id).
		Scan(&movimiento.IdMovimientoDinero, &movimiento.Nombre, &movimiento.Monto, &movimiento.EsIngreso, &movimiento.Activo, &movimiento.FechaCreacion, &movimiento.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "movimiento_ingreso_egreso no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar movimiento_ingreso_egreso")
		return
	}

	writeJSON(w, http.StatusOK, movimiento)
}

func CrearMovimientoIngresoEgreso(w http.ResponseWriter, r *http.Request) {
	var movimiento models.MovimientoIngresoEgreso
	if err := json.NewDecoder(r.Body).Decode(&movimiento); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	err := config.DB.QueryRow(`INSERT INTO movimiento_ingreso_egreso (nombre, monto, es_ingreso, activo) VALUES ($1, $2, $3, $4) RETURNING id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion`,
		movimiento.Nombre, movimiento.Monto, movimiento.EsIngreso, movimiento.Activo).
		Scan(&movimiento.IdMovimientoDinero, &movimiento.Nombre, &movimiento.Monto, &movimiento.EsIngreso, &movimiento.Activo, &movimiento.FechaCreacion, &movimiento.FechaModificacion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear movimiento_ingreso_egreso")
		return
	}

	writeJSON(w, http.StatusCreated, movimiento)
}

func ActualizarMovimientoIngresoEgreso(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var movimiento models.MovimientoIngresoEgreso
	if err := json.NewDecoder(r.Body).Decode(&movimiento); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	err = config.DB.QueryRow(`UPDATE movimiento_ingreso_egreso SET nombre = $1, monto = $2, es_ingreso = $3, activo = $4 WHERE id_movimiento_dinero = $5 RETURNING id_movimiento_dinero, nombre, monto, es_ingreso, activo, fecha_creacion, fecha_modificacion`,
		movimiento.Nombre, movimiento.Monto, movimiento.EsIngreso, movimiento.Activo, id).
		Scan(&movimiento.IdMovimientoDinero, &movimiento.Nombre, &movimiento.Monto, &movimiento.EsIngreso, &movimiento.Activo, &movimiento.FechaCreacion, &movimiento.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "movimiento_ingreso_egreso no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar movimiento_ingreso_egreso")
		return
	}

	writeJSON(w, http.StatusOK, movimiento)
}

func EliminarMovimientoIngresoEgreso(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM movimiento_ingreso_egreso WHERE id_movimiento_dinero = $1`, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar movimiento_ingreso_egreso porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar movimiento_ingreso_egreso")
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, "movimiento_ingreso_egreso no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "movimiento_ingreso_egreso eliminado correctamente"})
}
