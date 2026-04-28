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

func ObtenerLecciones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_leccion, id_modulo, id_contenido, titulo, descripcion, duracion_minutos, url_video, numero_leccion, activo, fecha_creacion, fecha_modificacion
		FROM leccion
		ORDER BY id_leccion
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar leccion")
		return
	}
	defer rows.Close()

	lecciones := []models.Leccion{}

	for rows.Next() {
		leccion, err := scanLeccion(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer leccion")
			return
		}
		lecciones = append(lecciones, leccion)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer leccion")
		return
	}

	writeJSON(w, http.StatusOK, lecciones)
}

func ObtenerLeccionPorID(w http.ResponseWriter, r *http.Request) {
	idLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_leccion, id_modulo, id_contenido, titulo, descripcion, duracion_minutos, url_video, numero_leccion, activo, fecha_creacion, fecha_modificacion
		FROM leccion
		WHERE id_leccion = $1
	`, idLeccion)

	leccion, err := scanLeccion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "leccion no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar leccion")
		return
	}

	writeJSON(w, http.StatusOK, leccion)
}

func CrearLeccion(w http.ResponseWriter, r *http.Request) {
	var leccion models.Leccion
	if err := json.NewDecoder(r.Body).Decode(&leccion); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesLeccion(leccion.IdModulo, leccion.IdContenido); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO leccion (id_modulo, id_contenido, titulo, descripcion, duracion_minutos, url_video, numero_leccion, activo)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id_leccion, id_modulo, id_contenido, titulo, descripcion, duracion_minutos, url_video, numero_leccion, activo, fecha_creacion, fecha_modificacion
	`, leccion.IdModulo, leccion.IdContenido, leccion.Titulo, leccion.Descripcion, leccion.DuracionMinutos, leccion.UrlVideo, leccion.NumeroLeccion, leccion.Activo)

	leccionCreada, err := scanLeccion(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear leccion")
		return
	}

	writeJSON(w, http.StatusCreated, leccionCreada)
}

func ActualizarLeccion(w http.ResponseWriter, r *http.Request) {
	idLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var leccion models.Leccion
	if err := json.NewDecoder(r.Body).Decode(&leccion); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if err := validarRelacionesLeccion(leccion.IdModulo, leccion.IdContenido); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	row := config.DB.QueryRow(`
		UPDATE leccion
		SET id_modulo = $1,
			id_contenido = $2,
			titulo = $3,
			descripcion = $4,
			duracion_minutos = $5,
			url_video = $6,
			numero_leccion = $7,
			activo = $8
		WHERE id_leccion = $9
		RETURNING id_leccion, id_modulo, id_contenido, titulo, descripcion, duracion_minutos, url_video, numero_leccion, activo, fecha_creacion, fecha_modificacion
	`, leccion.IdModulo, leccion.IdContenido, leccion.Titulo, leccion.Descripcion, leccion.DuracionMinutos, leccion.UrlVideo, leccion.NumeroLeccion, leccion.Activo, idLeccion)

	leccionActualizada, err := scanLeccion(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "leccion no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar leccion")
		return
	}

	writeJSON(w, http.StatusOK, leccionActualizada)
}

func EliminarLeccion(w http.ResponseWriter, r *http.Request) {
	idLeccion, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM leccion WHERE id_leccion = $1`, idLeccion)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar leccion porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar leccion")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de leccion")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "leccion no encontrada")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "leccion eliminada correctamente",
	})
}

func scanLeccion(scan func(dest ...any) error) (models.Leccion, error) {
	var leccion models.Leccion
	var idContenido sql.NullInt32
	var descripcion sql.NullString
	var duracionMinutos sql.NullInt32
	var urlVideo sql.NullString
	var numeroLeccion sql.NullInt32

	err := scan(
		&leccion.IdLeccion,
		&leccion.IdModulo,
		&idContenido,
		&leccion.Titulo,
		&descripcion,
		&duracionMinutos,
		&urlVideo,
		&numeroLeccion,
		&leccion.Activo,
		&leccion.FechaCreacion,
		&leccion.FechaModificacion,
	)
	if err != nil {
		return models.Leccion{}, err
	}

	leccion.IdContenido = nullInt32ToPointer(idContenido)
	leccion.Descripcion = nullStringToPointer(descripcion)
	leccion.DuracionMinutos = nullInt32ToPointer(duracionMinutos)
	leccion.UrlVideo = nullStringToPointer(urlVideo)
	leccion.NumeroLeccion = nullInt32ToPointer(numeroLeccion)

	return leccion, nil
}

func validarRelacionesLeccion(idModulo int, idContenido *int) error {
	existeModulo, err := existeRegistro("modulo_educativo", "id_modulo", idModulo)
	if err != nil {
		return errors.New("error al validar id_modulo")
	}
	if !existeModulo {
		return errors.New("id_modulo no existe")
	}

	if idContenido != nil {
		existeContenido, err := existeRegistro("contenido", "id_contenido", *idContenido)
		if err != nil {
			return errors.New("error al validar id_contenido")
		}
		if !existeContenido {
			return errors.New("id_contenido no existe")
		}
	}

	return nil
}



func existeRegistro(tabla string, columna string, id int) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM " + tabla + " WHERE " + columna + " = $1)"

	var existe bool
	err := config.DB.QueryRow(query, id).Scan(&existe)
	if err != nil {
		return false, err
	}



	return existe, nil
}
