package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"EDUCACION/config"
	"EDUCACION/models"

	"github.com/lib/pq"
)

func ObtenerProgresosLeccion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_progreso_leccion, id_usuario, id_leccion, completado, fecha_inicio, fecha_completado, activo, fecha_creacion, fecha_modificacion
		FROM progreso_leccion
		ORDER BY id_progreso_leccion
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar progreso_leccion")
		return
	}
	defer rows.Close()

	progresos := []models.ProgresoLeccion{}

	for rows.Next() {
		progreso, err := scanProgresoLeccion(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer progreso_leccion")
			return
		}
		progresos = append(progresos, progreso)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer progreso_leccion")
		return
	}

	writeJSON(w, http.StatusOK, progresos)
}

func ObtenerProgresoLeccionPorID(w http.ResponseWriter, r *http.Request) {
	idProgresoLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_progreso_leccion, id_usuario, id_leccion, completado, fecha_inicio, fecha_completado, activo, fecha_creacion, fecha_modificacion
		FROM progreso_leccion
		WHERE id_progreso_leccion = $1
	`, idProgresoLeccion)

	progreso, err := scanProgresoLeccion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "progreso_leccion no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar progreso_leccion")
		return
	}

	writeJSON(w, http.StatusOK, progreso)
}

func CrearProgresoLeccion(w http.ResponseWriter, r *http.Request) {
	var progreso models.ProgresoLeccion
	if err := json.NewDecoder(r.Body).Decode(&progreso); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesProgresoLeccion(progreso.IdUsuario, progreso.IdLeccion); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO progreso_leccion (id_usuario, id_leccion, completado, fecha_inicio, fecha_completado, activo)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_progreso_leccion, id_usuario, id_leccion, completado, fecha_inicio, fecha_completado, activo, fecha_creacion, fecha_modificacion
	`, progreso.IdUsuario, progreso.IdLeccion, progreso.Completado, progreso.FechaInicio, progreso.FechaCompletado, progreso.Activo)

	progresoCreado, err := scanProgresoLeccion(row.Scan)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un progreso_leccion para ese id_usuario e id_leccion")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear progreso_leccion")
		return
	}

	writeJSON(w, http.StatusCreated, progresoCreado)
}

func ActualizarProgresoLeccion(w http.ResponseWriter, r *http.Request) {
	idProgresoLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var progreso models.ProgresoLeccion
	if err := json.NewDecoder(r.Body).Decode(&progreso); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesProgresoLeccion(progreso.IdUsuario, progreso.IdLeccion); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		UPDATE progreso_leccion
		SET id_usuario = $1,
			id_leccion = $2,
			completado = $3,
			fecha_inicio = $4,
			fecha_completado = $5,
			activo = $6
		WHERE id_progreso_leccion = $7
		RETURNING id_progreso_leccion, id_usuario, id_leccion, completado, fecha_inicio, fecha_completado, activo, fecha_creacion, fecha_modificacion
	`, progreso.IdUsuario, progreso.IdLeccion, progreso.Completado, progreso.FechaInicio, progreso.FechaCompletado, progreso.Activo, idProgresoLeccion)

	progresoActualizado, err := scanProgresoLeccion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "progreso_leccion no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			writeError(w, http.StatusConflict, "ya existe un progreso_leccion para ese id_usuario e id_leccion")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar progreso_leccion")
		return
	}

	writeJSON(w, http.StatusOK, progresoActualizado)
}

func EliminarProgresoLeccion(w http.ResponseWriter, r *http.Request) {
	idProgresoLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM progreso_leccion WHERE id_progreso_leccion = $1`, idProgresoLeccion)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al eliminar progreso_leccion")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de progreso_leccion")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "progreso_leccion no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "progreso_leccion eliminado correctamente",
	})
}

func scanProgresoLeccion(scan func(dest ...any) error) (models.ProgresoLeccion, error) {
	var progreso models.ProgresoLeccion
	var completado sql.NullBool
	var fechaInicio sql.NullTime
	var fechaCompletado sql.NullTime

	err := scan(
		&progreso.IdProgresoLeccion,
		&progreso.IdUsuario,
		&progreso.IdLeccion,
		&completado,
		&fechaInicio,
		&fechaCompletado,
		&progreso.Activo,
		&progreso.FechaCreacion,
		&progreso.FechaModificacion,
	)
	if err != nil {
		return models.ProgresoLeccion{}, err
	}

	progreso.Completado = nullBoolToPointer(completado)
	progreso.FechaInicio = nullTimeToPointer(fechaInicio)
	progreso.FechaCompletado = nullTimeToPointer(fechaCompletado)

	return progreso, nil
}

func validarRelacionesProgresoLeccion(idUsuario int, idLeccion int) error {
	existeUsuario, err := existeRegistro("auth.usuario", "id_usuario", idUsuario)
	if err != nil {
		return errors.New("error al validar id_usuario")
	}
	if !existeUsuario {
		return errors.New("id_usuario no existe")
	}

	existeLeccion, err := existeRegistro("leccion", "id_leccion", idLeccion)
	if err != nil {
		return errors.New("error al validar id_leccion")
	}
	if !existeLeccion {
		return errors.New("id_leccion no existe")
	}

	return nil
}

func nullBoolToPointer(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}

	result := value.Bool
	return &result
}
