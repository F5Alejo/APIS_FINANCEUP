package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/models"
)

func ObtenerTiposIngresoInversion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_tipo_ingreso, id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo, fecha_creacion, fecha_modificacion FROM tipo_ingreso_inversion ORDER BY id_tipo_ingreso`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar tipo_ingreso_inversion")
		return
	}
	defer rows.Close()
	items := []models.TipoIngresoInversion{}
	for rows.Next() {
		item, err := scanTipoIngresoInversion(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer tipo_ingreso_inversion")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerTipoIngresoInversionPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_tipo_ingreso, id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo, fecha_creacion, fecha_modificacion FROM tipo_ingreso_inversion WHERE id_tipo_ingreso = $1`, id)
	item, err := scanTipoIngresoInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tipo_ingreso_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar tipo_ingreso_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearTipoIngresoInversion(w http.ResponseWriter, r *http.Request) {
	var item models.TipoIngresoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarMovimientoInversionOpcional(item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`INSERT INTO tipo_ingreso_inversion (id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo) VALUES ($1, $2, $3, $4) RETURNING id_tipo_ingreso, id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.IdMovimientoDinero, item.NombreMovimientoPago, item.Descripcion, item.Activo)
	item, err := scanTipoIngresoInversion(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear tipo_ingreso_inversion")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarTipoIngresoInversion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TipoIngresoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarMovimientoInversionOpcional(item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`UPDATE tipo_ingreso_inversion SET id_movimiento_dinero = $1, nombre_movimiento_pago = $2, descripcion = $3, activo = $4 WHERE id_tipo_ingreso = $5 RETURNING id_tipo_ingreso, id_movimiento_dinero, nombre_movimiento_pago, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.IdMovimientoDinero, item.NombreMovimientoPago, item.Descripcion, item.Activo, id)
	item, err = scanTipoIngresoInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tipo_ingreso_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar tipo_ingreso_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarTipoIngresoInversion(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "tipo_ingreso_inversion", "id_tipo_ingreso", "tipo_ingreso_inversion")
}

func scanTipoIngresoInversion(scan func(dest ...any) error) (models.TipoIngresoInversion, error) {
	var item models.TipoIngresoInversion
	var idMovimiento sql.NullInt32
	var descripcion sql.NullString
	err := scan(&item.IdTipoIngreso, &idMovimiento, &item.NombreMovimientoPago, &descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.TipoIngresoInversion{}, err
	}
	item.IdMovimientoDinero = nullInt32ToPointer(idMovimiento)
	item.Descripcion = nullStringToPointer(descripcion)
	return item, nil
}

func validarMovimientoInversionOpcional(id *int) error {
	if id == nil {
		return nil
	}
	existe, err := existeRegistro("movimiento_inversion", "id_movimiento_dinero", *id)
	if err != nil {
		return errors.New("error al validar id_movimiento_dinero")
	}
	if !existe {
		return errors.New("id_movimiento_dinero no existe")
	}
	return nil
}
