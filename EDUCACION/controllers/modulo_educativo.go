package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"EDUCACION/config"
	"EDUCACION/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func ObtenerModulosEducativos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_modulo, titulo, descripcion, contenido, nivel, url_thumbnail, activo, fecha_creacion, fecha_modificacion
		FROM modulo_educativo
		ORDER BY id_modulo
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar modulo_educativo")
		return
	}
	defer rows.Close()

	modulos := []models.ModuloEducativo{}

	for rows.Next() {
		modulo, err := scanModuloEducativo(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer modulo_educativo")
			return
		}
		modulos = append(modulos, modulo)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer modulo_educativo")
		return
	}

	writeJSON(w, http.StatusOK, modulos)
}

func ObtenerModuloEducativoPorID(w http.ResponseWriter, r *http.Request) {
	idModulo, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_modulo, titulo, descripcion, contenido, nivel, url_thumbnail, activo, fecha_creacion, fecha_modificacion
		FROM modulo_educativo
		WHERE id_modulo = $1
	`, idModulo)

	modulo, err := scanModuloEducativo(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "modulo_educativo no encontrado")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar modulo_educativo")
		return
	}

	writeJSON(w, http.StatusOK, modulo)
}

func CrearModuloEducativo(w http.ResponseWriter, r *http.Request) {
	var modulo models.ModuloEducativo
	if err := json.NewDecoder(r.Body).Decode(&modulo); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO modulo_educativo (titulo, descripcion, contenido, nivel, url_thumbnail, activo)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_modulo, titulo, descripcion, contenido, nivel, url_thumbnail, activo, fecha_creacion, fecha_modificacion
	`, modulo.Titulo, modulo.Descripcion, modulo.Contenido, modulo.Nivel, modulo.UrlThumbnail, modulo.Activo)

	createdModulo, err := scanModuloEducativo(row.Scan)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "nivel no valido para modulo_educativo")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al crear modulo_educativo")
		return
	}

	writeJSON(w, http.StatusCreated, createdModulo)
}

func ActualizarModuloEducativo(w http.ResponseWriter, r *http.Request) {
	idModulo, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var modulo models.ModuloEducativo
	if err := json.NewDecoder(r.Body).Decode(&modulo); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		UPDATE modulo_educativo
		SET titulo = $1,
			descripcion = $2,
			contenido = $3,
			nivel = $4,
			url_thumbnail = $5,
			activo = $6
		WHERE id_modulo = $7
		RETURNING id_modulo, titulo, descripcion, contenido, nivel, url_thumbnail, activo, fecha_creacion, fecha_modificacion
	`, modulo.Titulo, modulo.Descripcion, modulo.Contenido, modulo.Nivel, modulo.UrlThumbnail, modulo.Activo, idModulo)

	updatedModulo, err := scanModuloEducativo(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "modulo_educativo no encontrado")
			return
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			writeError(w, http.StatusBadRequest, "nivel no valido para modulo_educativo")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar modulo_educativo")
		return
	}

	writeJSON(w, http.StatusOK, updatedModulo)
}

func EliminarModuloEducativo(w http.ResponseWriter, r *http.Request) {
	idModulo, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM modulo_educativo WHERE id_modulo = $1`, idModulo)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar modulo_educativo porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar modulo_educativo")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de modulo_educativo")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "modulo_educativo no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "modulo_educativo eliminado correctamente",
	})
}

func scanModuloEducativo(scan func(dest ...any) error) (models.ModuloEducativo, error) {
	var modulo models.ModuloEducativo
	var descripcion sql.NullString
	var contenido sql.NullString
	var urlThumbnail sql.NullString

	err := scan(
		&modulo.IdModulo,
		&modulo.Titulo,
		&descripcion,
		&contenido,
		&modulo.Nivel,
		&urlThumbnail,
		&modulo.Activo,
		&modulo.FechaCreacion,
		&modulo.FechaModificacion,
	)
	if err != nil {
		return models.ModuloEducativo{}, err
	}

	modulo.Descripcion = nullStringToPointer(descripcion)
	modulo.Contenido = nullStringToPointer(contenido)
	modulo.UrlThumbnail = nullStringToPointer(urlThumbnail)

	return modulo, nil
}

func nullStringToPointer(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	text := value.String
	return &text
}

func getIDFromRequest(r *http.Request, key string) (int, error) {
	idParam := mux.Vars(r)[key]
	return strconv.Atoi(idParam)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{
		"error": message,
	})
}
