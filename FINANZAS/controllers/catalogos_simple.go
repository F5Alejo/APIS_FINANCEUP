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

func ObtenerTiposInversion(w http.ResponseWriter, r *http.Request) { listarTipoInversion(w) }
func ObtenerTipoInversionPorID(w http.ResponseWriter, r *http.Request) {
	obtenerTipoInversionPorID(w, r)
}
func CrearTipoInversion(w http.ResponseWriter, r *http.Request)      { crearTipoInversion(w, r) }
func ActualizarTipoInversion(w http.ResponseWriter, r *http.Request) { actualizarTipoInversion(w, r) }
func EliminarTipoInversion(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "tipo_inversion", "id_tipo_inversion", "tipo_inversion")
}

func ObtenerNivelesRiesgo(w http.ResponseWriter, r *http.Request)    { listarNivelRiesgo(w) }
func ObtenerNivelRiesgoPorID(w http.ResponseWriter, r *http.Request) { obtenerNivelRiesgoPorID(w, r) }
func CrearNivelRiesgo(w http.ResponseWriter, r *http.Request)        { crearNivelRiesgo(w, r) }
func ActualizarNivelRiesgo(w http.ResponseWriter, r *http.Request)   { actualizarNivelRiesgo(w, r) }
func EliminarNivelRiesgo(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "nivel_riesgo", "id_nivel_riesgo", "nivel_riesgo")
}

func ObtenerEditarMetas(w http.ResponseWriter, r *http.Request)     { listarEditarMeta(w) }
func ObtenerEditarMetaPorID(w http.ResponseWriter, r *http.Request) { obtenerEditarMetaPorID(w, r) }
func CrearEditarMeta(w http.ResponseWriter, r *http.Request)        { crearEditarMeta(w, r) }
func ActualizarEditarMeta(w http.ResponseWriter, r *http.Request)   { actualizarEditarMeta(w, r) }
func EliminarEditarMeta(w http.ResponseWriter, r *http.Request) {
	eliminarGenerico(w, r, "editar_meta", "id_editar_meta", "editar_meta")
}

func listarTipoInversion(w http.ResponseWriter) {
	rows, err := config.DB.Query(`SELECT id_tipo_inversion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion FROM tipo_inversion ORDER BY id_tipo_inversion`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar tipo_inversion")
		return
	}
	defer rows.Close()
	items := []models.TipoInversion{}
	for rows.Next() {
		item, err := scanTipoInversion(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer tipo_inversion")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func obtenerTipoInversionPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_tipo_inversion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion FROM tipo_inversion WHERE id_tipo_inversion = $1`, id)
	item, err := scanTipoInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tipo_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar tipo_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func crearTipoInversion(w http.ResponseWriter, r *http.Request) {
	var item models.TipoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO tipo_inversion (nombre, descripcion, activo) VALUES ($1, $2, $3) RETURNING id_tipo_inversion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion`, item.Nombre, item.Descripcion, item.Activo)
	item, err := scanTipoInversion(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear tipo_inversion")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func actualizarTipoInversion(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.TipoInversion
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`UPDATE tipo_inversion SET nombre = $1, descripcion = $2, activo = $3 WHERE id_tipo_inversion = $4 RETURNING id_tipo_inversion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion`, item.Nombre, item.Descripcion, item.Activo, id)
	item, err = scanTipoInversion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "tipo_inversion no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar tipo_inversion")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func scanTipoInversion(scan func(dest ...any) error) (models.TipoInversion, error) {
	var item models.TipoInversion
	var descripcion sql.NullString
	err := scan(&item.IdTipoInversion, &item.Nombre, &descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.TipoInversion{}, err
	}
	item.Descripcion = nullStringToPointer(descripcion)
	return item, nil
}

func listarNivelRiesgo(w http.ResponseWriter) {
	rows, err := config.DB.Query(`SELECT id_nivel_riesgo, nombre, activo, fecha_creacion, fecha_modificacion FROM nivel_riesgo ORDER BY id_nivel_riesgo`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar nivel_riesgo")
		return
	}
	defer rows.Close()
	items := []models.NivelRiesgo{}
	for rows.Next() {
		var item models.NivelRiesgo
		if err := rows.Scan(&item.IdNivelRiesgo, &item.Nombre, &item.Activo, &item.FechaCreacion, &item.FechaModificacion); err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer nivel_riesgo")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func obtenerNivelRiesgoPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.NivelRiesgo
	err = config.DB.QueryRow(`SELECT id_nivel_riesgo, nombre, activo, fecha_creacion, fecha_modificacion FROM nivel_riesgo WHERE id_nivel_riesgo = $1`, id).Scan(&item.IdNivelRiesgo, &item.Nombre, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "nivel_riesgo no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar nivel_riesgo")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func crearNivelRiesgo(w http.ResponseWriter, r *http.Request) {
	var item models.NivelRiesgo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err := config.DB.QueryRow(`INSERT INTO nivel_riesgo (nombre, activo) VALUES ($1, $2) RETURNING id_nivel_riesgo, nombre, activo, fecha_creacion, fecha_modificacion`, item.Nombre, item.Activo).Scan(&item.IdNivelRiesgo, &item.Nombre, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear nivel_riesgo")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func actualizarNivelRiesgo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.NivelRiesgo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	err = config.DB.QueryRow(`UPDATE nivel_riesgo SET nombre = $1, activo = $2 WHERE id_nivel_riesgo = $3 RETURNING id_nivel_riesgo, nombre, activo, fecha_creacion, fecha_modificacion`, item.Nombre, item.Activo, id).Scan(&item.IdNivelRiesgo, &item.Nombre, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "nivel_riesgo no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar nivel_riesgo")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func listarEditarMeta(w http.ResponseWriter) {
	rows, err := config.DB.Query(`SELECT id_editar_meta, nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo, fecha_creacion, fecha_modificacion FROM editar_meta ORDER BY id_editar_meta`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar editar_meta")
		return
	}
	defer rows.Close()
	items := []models.EditarMeta{}
	for rows.Next() {
		item, err := scanEditarMeta(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer editar_meta")
			return
		}
		items = append(items, item)
	}
	writeJSON(w, http.StatusOK, items)
}

func obtenerEditarMetaPorID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	row := config.DB.QueryRow(`SELECT id_editar_meta, nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo, fecha_creacion, fecha_modificacion FROM editar_meta WHERE id_editar_meta = $1`, id)
	item, err := scanEditarMeta(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "editar_meta no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al consultar editar_meta")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func crearEditarMeta(w http.ResponseWriter, r *http.Request) {
	var item models.EditarMeta
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`INSERT INTO editar_meta (nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_editar_meta, nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.Nombre, item.MontoActual, item.MontoObjetivo, item.AhorroMensual, item.FechaObjetivo, item.Descripcion, item.Activo)
	item, err := scanEditarMeta(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear editar_meta")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func actualizarEditarMeta(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	var item models.EditarMeta
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	row := config.DB.QueryRow(`UPDATE editar_meta SET nombre = $1, monto_actual = $2, monto_objetivo = $3, ahorro_mensual = $4, fecha_objetivo = $5, descripcion = $6, activo = $7 WHERE id_editar_meta = $8 RETURNING id_editar_meta, nombre, monto_actual, monto_objetivo, ahorro_mensual, fecha_objetivo, descripcion, activo, fecha_creacion, fecha_modificacion`,
		item.Nombre, item.MontoActual, item.MontoObjetivo, item.AhorroMensual, item.FechaObjetivo, item.Descripcion, item.Activo, id)
	item, err = scanEditarMeta(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "editar_meta no encontrado")
		} else {
			writeError(w, http.StatusInternalServerError, "error al actualizar editar_meta")
		}
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func scanEditarMeta(scan func(dest ...any) error) (models.EditarMeta, error) {
	var item models.EditarMeta
	var nombre sql.NullString
	var montoActual sql.NullFloat64
	var ahorroMensual sql.NullFloat64
	var fechaObjetivo sql.NullTime
	var descripcion sql.NullString
	err := scan(&item.IdEditarMeta, &nombre, &montoActual, &item.MontoObjetivo, &ahorroMensual, &fechaObjetivo, &descripcion, &item.Activo, &item.FechaCreacion, &item.FechaModificacion)
	if err != nil {
		return models.EditarMeta{}, err
	}
	item.Nombre = nullStringToPointer(nombre)
	item.MontoActual = nullFloat64ToPointer(montoActual)
	item.AhorroMensual = nullFloat64ToPointer(ahorroMensual)
	item.FechaObjetivo = nullTimeToPointer(fechaObjetivo)
	item.Descripcion = nullStringToPointer(descripcion)
	return item, nil
}

func eliminarGenerico(w http.ResponseWriter, r *http.Request, tabla string, columna string, nombre string) {
	id, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}
	result, err := config.DB.Exec("DELETE FROM "+tabla+" WHERE "+columna+" = $1", id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar "+nombre+" porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar "+nombre)
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		writeError(w, http.StatusNotFound, nombre+" no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": nombre + " eliminado correctamente"})
}
