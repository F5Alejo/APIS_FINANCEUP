package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"EDUCACION/config"
	"EDUCACION/models"

	"github.com/lib/pq"
)

func ObtenerProgresosEducativos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_progreso, id_usuario, id_modulo, porcentaje_completado, fecha_inicio, fecha_completado, calificacion, activo, fecha_creacion, fecha_modificacion
		FROM progreso_educativo
		ORDER BY id_progreso
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar progreso_educativo")
		return
	}
	defer rows.Close()

	progresos := []models.ProgresoEducativo{}

	for rows.Next() {
		progreso, err := scanProgresoEducativo(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer progreso_educativo")
			return
		}
		progresos = append(progresos, progreso)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer progreso_educativo")
		return
	}

	writeJSON(w, http.StatusOK, progresos)
}

func ObtenerProgresoEducativoPorID(w http.ResponseWriter, r *http.Request) {
	idProgreso, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_progreso, id_usuario, id_modulo, porcentaje_completado, fecha_inicio, fecha_completado, calificacion, activo, fecha_creacion, fecha_modificacion
		FROM progreso_educativo
		WHERE id_progreso = $1
	`, idProgreso)

	progreso, err := scanProgresoEducativo(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "progreso_educativo no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar progreso_educativo")
		return
	}

	writeJSON(w, http.StatusOK, progreso)
}

func CrearProgresoEducativo(w http.ResponseWriter, r *http.Request) {
	var progreso models.ProgresoEducativo
	if err := json.NewDecoder(r.Body).Decode(&progreso); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesProgresoEducativo(progreso.IdUsuario, progreso.IdModulo); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO progreso_educativo (id_usuario, id_modulo, porcentaje_completado, fecha_inicio, fecha_completado, calificacion, activo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_progreso, id_usuario, id_modulo, porcentaje_completado, fecha_inicio, fecha_completado, calificacion, activo, fecha_creacion, fecha_modificacion
	`, progreso.IdUsuario, progreso.IdModulo, progreso.PorcentajeCompletado, progreso.FechaInicio, progreso.FechaCompletado, progreso.Calificacion, progreso.Activo)

	progresoCreado, err := scanProgresoEducativo(row.Scan)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un progreso_educativo para ese id_usuario e id_modulo")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear progreso_educativo")
		return
	}

	writeJSON(w, http.StatusCreated, progresoCreado)
}

func ActualizarProgresoEducativo(w http.ResponseWriter, r *http.Request) {
	idProgreso, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var progreso models.ProgresoEducativo
	if err := json.NewDecoder(r.Body).Decode(&progreso); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesProgresoEducativo(progreso.IdUsuario, progreso.IdModulo); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		UPDATE progreso_educativo
		SET id_usuario = $1,
			id_modulo = $2,
			porcentaje_completado = $3,
			fecha_inicio = $4,
			fecha_completado = $5,
			calificacion = $6,
			activo = $7
		WHERE id_progreso = $8
		RETURNING id_progreso, id_usuario, id_modulo, porcentaje_completado, fecha_inicio, fecha_completado, calificacion, activo, fecha_creacion, fecha_modificacion
	`, progreso.IdUsuario, progreso.IdModulo, progreso.PorcentajeCompletado, progreso.FechaInicio, progreso.FechaCompletado, progreso.Calificacion, progreso.Activo, idProgreso)

	progresoActualizado, err := scanProgresoEducativo(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "progreso_educativo no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un progreso_educativo para ese id_usuario e id_modulo")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar progreso_educativo")
		return
	}

	writeJSON(w, http.StatusOK, progresoActualizado)
}

func EliminarProgresoEducativo(w http.ResponseWriter, r *http.Request) {
	idProgreso, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM progreso_educativo WHERE id_progreso = $1`, idProgreso)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al eliminar progreso_educativo")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de progreso_educativo")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "progreso_educativo no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "progreso_educativo eliminado correctamente",
	})
}

func scanProgresoEducativo(scan func(dest ...any) error) (models.ProgresoEducativo, error) {
	var progreso models.ProgresoEducativo
	var porcentajeCompletado sql.NullInt32
	var fechaInicio sql.NullTime
	var fechaCompletado sql.NullTime
	var calificacion sql.NullInt32

	err := scan(
		&progreso.IdProgreso,
		&progreso.IdUsuario,
		&progreso.IdModulo,
		&porcentajeCompletado,
		&fechaInicio,
		&fechaCompletado,
		&calificacion,
		&progreso.Activo,
		&progreso.FechaCreacion,
		&progreso.FechaModificacion,
	)
	if err != nil {
		return models.ProgresoEducativo{}, err
	}

	progreso.PorcentajeCompletado = nullInt32ToPointer(porcentajeCompletado)
	progreso.FechaInicio = nullTimeToPointer(fechaInicio)
	progreso.FechaCompletado = nullTimeToPointer(fechaCompletado)
	progreso.Calificacion = nullInt32ToPointer(calificacion)

	return progreso, nil
}

func validarRelacionesProgresoEducativo(idUsuario int, idModulo int) error {
	existeUsuario, err := existeRegistro("auth.usuario", "id_usuario", idUsuario)
	if err != nil {
		return errors.New("error al validar id_usuario")
	}
	if !existeUsuario {
		return errors.New("id_usuario no existe")
	}

	existeModulo, err := existeRegistro("modulo_educativo", "id_modulo", idModulo)
	if err != nil {
		return errors.New("error al validar id_modulo")
	}
	if !existeModulo {
		return errors.New("id_modulo no existe")
	}

	return nil
}

func nullTimeToPointer(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	fecha := value.Time
	return &fecha
}
