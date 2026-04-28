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

func ObtenerContenidos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_contenido, titulo, descripcion, duracion_minutos, url_video, activo, fecha_creacion, fecha_modificacion
		FROM contenido
		ORDER BY id_contenido
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar contenido")
		return
	}
	defer rows.Close()

	contenidos := []models.Contenido{}

	for rows.Next() {
		contenido, err := scanContenido(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer contenido")
			return
		}
		contenidos = append(contenidos, contenido)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer contenido")
		return
	}

	writeJSON(w, http.StatusOK, contenidos)
}

func ObtenerContenidoPorID(w http.ResponseWriter, r *http.Request) {
	idContenido, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_contenido, titulo, descripcion, duracion_minutos, url_video, activo, fecha_creacion, fecha_modificacion
		FROM contenido
		WHERE id_contenido = $1
	`, idContenido)

	contenido, err := scanContenido(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "contenido no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar contenido")
		return
	}

	writeJSON(w, http.StatusOK, contenido)
}

func CrearContenido(w http.ResponseWriter, r *http.Request) {
	var contenido models.Contenido
	if err := json.NewDecoder(r.Body).Decode(&contenido); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO contenido (titulo, descripcion, duracion_minutos, url_video, activo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id_contenido, titulo, descripcion, duracion_minutos, url_video, activo, fecha_creacion, fecha_modificacion
	`, contenido.Titulo, contenido.Descripcion, contenido.DuracionMinutos, contenido.UrlVideo, contenido.Activo)

	createdContenido, err := scanContenido(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear contenido")
		return
	}

	writeJSON(w, http.StatusCreated, createdContenido)
}

func ActualizarContenido(w http.ResponseWriter, r *http.Request) {
	idContenido, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var contenido models.Contenido
	if err := json.NewDecoder(r.Body).Decode(&contenido); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		UPDATE contenido
		SET titulo = $1,
			descripcion = $2,
			duracion_minutos = $3,
			url_video = $4,
			activo = $5
		WHERE id_contenido = $6
		RETURNING id_contenido, titulo, descripcion, duracion_minutos, url_video, activo, fecha_creacion, fecha_modificacion
	`, contenido.Titulo, contenido.Descripcion, contenido.DuracionMinutos, contenido.UrlVideo, contenido.Activo, idContenido)

	updatedContenido, err := scanContenido(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "contenido no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar contenido")
		return
	}

	writeJSON(w, http.StatusOK, updatedContenido)
}

func EliminarContenido(w http.ResponseWriter, r *http.Request) {
	idContenido, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM contenido WHERE id_contenido = $1`, idContenido)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar contenido porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar contenido")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de contenido")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "contenido no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "contenido eliminado correctamente",
	})
}

func scanContenido(scan func(dest ...any) error) (models.Contenido, error) {
	var contenido models.Contenido
	var descripcion sql.NullString
	var duracionMinutos sql.NullInt32
	var urlVideo sql.NullString

	err := scan(
		&contenido.IdContenido,
		&contenido.Titulo,
		&descripcion,
		&duracionMinutos,
		&urlVideo,
		&contenido.Activo,
		&contenido.FechaCreacion,
		&contenido.FechaModificacion,
	)
	if err != nil {
		return models.Contenido{}, err
	}

	contenido.Descripcion = nullStringToPointer(descripcion)
	contenido.DuracionMinutos = nullInt32ToPointer(duracionMinutos)
	contenido.UrlVideo = nullStringToPointer(urlVideo)

	return contenido, nil
}

func nullInt32ToPointer(value sql.NullInt32) *int {
	if !value.Valid {
		return nil
	}

	number := int(value.Int32)
	return &number
}
