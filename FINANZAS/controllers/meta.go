package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"FINANZAS/config"
	"FINANZAS/models"
)

func ObtenerMetas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id_meta, id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo, fecha_creacion, fecha_modificacion FROM meta ORDER BY id_meta`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar meta")
		return
	}
	defer rows.Close()
	items := []models.Meta{}
	for rows.Next() {
		item, err := scanMeta(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer meta")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func ObtenerMetaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_meta, id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo, fecha_creacion, fecha_modificacion FROM meta WHERE id_meta = $1`, id)
	item, err := scanMeta(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "meta no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar meta")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func CrearMeta(w http.ResponseWriter, r *http.Request) {
	var item models.Meta
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesMeta(item.IdUsuario, item.IdEditarMeta, item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`INSERT INTO meta (id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id_meta, id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdEditarMeta, item.IdMovimientoDinero, item.Nombre, item.Descripcion, item.MontoObjetivo, item.MontoActual, item.FechaLimite, item.Color, item.Activo)
	item, err := scanMeta(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear meta")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func ActualizarMeta(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.Meta
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if err := validarRelacionesMeta(item.IdUsuario, item.IdEditarMeta, item.IdMovimientoDinero); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	row := config.DB.QueryRow(`UPDATE meta SET id_usuario = $1, id_editar_meta = $2, id_movimiento_dinero = $3, nombre = $4, descripcion = $5, monto_objetivo = $6, monto_actual = $7, fecha_limite = $8, color = $9, activo = $10 WHERE id_meta = $11 RETURNING id_meta, id_usuario, id_editar_meta, id_movimiento_dinero, nombre, descripcion, monto_objetivo, monto_actual, fecha_limite, color, activo, fecha_creacion, fecha_modificacion`,
		item.IdUsuario, item.IdEditarMeta, item.IdMovimientoDinero, item.Nombre, item.Descripcion, item.MontoObjetivo, item.MontoActual, item.FechaLimite, item.Color, item.Activo, id)
	item, err = scanMeta(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "meta no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar meta")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func EliminarMeta(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "meta", "id_meta", "meta")
}

func scanMeta(scan func(dest ...any) error) (models.Meta, error) {
	var item models.Meta
	var idMovimiento sql.NullInt32
	var nombre sql.NullString
	var descripcion sql.NullString
	var montoObjetivo sql.NullFloat64
	var montoActual sql.NullFloat64
	var fechaLimite sql.NullTime
	var color sql.NullString
	err := scan(&item.IdMeta, &item.IdUsuario, &item.IdEditarMeta, &idMovimiento, &nombre, &descripcion, &montoObjetivo, &montoActual, &fechaLimite, &color, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.Meta{}, err
	}
	item.IdMovimientoDinero = nullInt32ToPointer(idMovimiento)
	item.Nombre = nullStringToPointer(nombre)
	item.Descripcion = nullStringToPointer(descripcion)
	item.MontoObjetivo = nullFloat64ToPointer(montoObjetivo)
	item.MontoActual = nullFloat64ToPointer(montoActual)
	item.FechaLimite = nullTimeToPointer(fechaLimite)
	item.Color = nullStringToPointer(color)
	return item, nil
}

func validarRelacionesMeta(idUsuario int, idEditarMeta int, idMovimiento *int) error {
	if ok, err := existeRegistro("auth.usuario", "id_usuario", idUsuario); err != nil {
		return errors.New("error al validar id_usuario")
	} else if !ok {
		return errors.New("id_usuario no existe")
	}
	if ok, err := existeRegistro("editar_meta", "id_editar_meta", idEditarMeta); err != nil {
		return errors.New("error al validar id_editar_meta")
	} else if !ok {
		return errors.New("id_editar_meta no existe")
	}
	if idMovimiento != nil {
		if ok, err := existeRegistro("movimiento_meta", "id_movimiento_dinero", *idMovimiento); err != nil {
			return errors.New("error al validar id_movimiento_dinero")
		} else if !ok {
			return errors.New("id_movimiento_dinero no existe")
		}
	}
	return nil
}
