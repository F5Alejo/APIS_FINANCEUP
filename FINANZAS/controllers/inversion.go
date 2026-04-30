package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/models"
)

func ObtenerInversiones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_inversion, id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo, fecha_creacion, fecha_modificacion FROM inversion ORDER BY id_inversion`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar inversion")
		return
	}
	defer rows.Close()
	items := []models.Inversion{}
	for rows.Next() {
		item, err := scanInversion(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer inversion")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerInversionPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_inversion, id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo, fecha_creacion, fecha_modificacion FROM inversion WHERE id_inversion = $1`, id)
	item, err := scanInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearInversion(w http.ResponseWriter, r *http.Request) {
	var item models.Inversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesInversion(item.IdUsuario, item.IdTipoInversion, item.IdNivelRiesgo, item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`INSERT INTO inversion (id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id_inversion, id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdTipoInversion, item.IdNivelRiesgo, item.IdMovimientoDinero, item.Nombre, item.Monto, item.Rentabilidad, item.FechaInicio, item.FechaFin, item.Activo)
	item, err := scanInversion(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear inversion")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarInversion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Inversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesInversion(item.IdUsuario, item.IdTipoInversion, item.IdNivelRiesgo, item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`UPDATE inversion SET id_usuario = $1, id_tipo_inversion = $2, id_nivel_riesgo = $3, id_movimiento_dinero = $4, nombre = $5, monto = $6, rentabilidad = $7, fecha_inicio = $8, fecha_fin = $9, activo = $10 WHERE id_inversion = $11 RETURNING id_inversion, id_usuario, id_tipo_inversion, id_nivel_riesgo, id_movimiento_dinero, nombre, monto, rentabilidad, fecha_inicio, fecha_fin, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdTipoInversion, item.IdNivelRiesgo, item.IdMovimientoDinero, item.Nombre, item.Monto, item.Rentabilidad, item.FechaInicio, item.FechaFin, item.Activo, id)
	item, err = scanInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarInversion(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "inversion", "id_inversion", "inversion")
}

func scanInversion(scan func(dest ...any) error) (models.Inversion, error) {
	var item models.Inversion
	var idMovimiento sql.NullInt32
	var nombre sql.NullString
	var monto sql.NullFloat64
	var rentabilidad sql.NullFloat64
	var fechaInicio sql.NullTime
	var fechaFin sql.NullTime
	err := scan(&item.IdInversion, &item.IdUsuario, &item.IdTipoInversion, &item.IdNivelRiesgo, &idMovimiento, &nombre, &monto, &rentabilidad, &fechaInicio, &fechaFin, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.Inversion{}, err
	}
	item.IdMovimientoDinero = nullInt32ToPointer(idMovimiento)
	item.Nombre = nullStringToPointer(nombre)
	item.Monto = nullFloat64ToPointer(monto)
	item.Rentabilidad = nullFloat64ToPointer(rentabilidad)
	item.FechaInicio = nullTimeToPointer(fechaInicio)
	item.FechaFin = nullTimeToPointer(fechaFin)
	return item, nil
}

func validarRelacionesInversion(idUsuario int, idTipoInversion int, idNivelRiesgo int, idMovimiento *int) error {
	if ok, err := existeRegistro("auth.usuario", "id_usuario", idUsuario); err != nil {
		return errors.New("error al validar id_usuario")
	} else if !ok {
		return errors.New("id_usuario no existe")
	}
	if ok, err := existeRegistro("tipo_inversion", "id_tipo_inversion", idTipoInversion); err != nil {
		return errors.New("error al validar id_tipo_inversion")
	} else if !ok {
		return errors.New("id_tipo_inversion no existe")
	}
	if ok, err := existeRegistro("nivel_riesgo", "id_nivel_riesgo", idNivelRiesgo); err != nil {
		return errors.New("error al validar id_nivel_riesgo")
	} else if !ok {
		return errors.New("id_nivel_riesgo no existe")
	}
	if idMovimiento != nil {
		if ok, err := existeRegistro("movimiento_inversion", "id_movimiento_dinero", *idMovimiento); err != nil {
			return errors.New("error al validar id_movimiento_dinero")
		} else if !ok {
			return errors.New("id_movimiento_dinero no existe")
		}
	}
	return nil
}
