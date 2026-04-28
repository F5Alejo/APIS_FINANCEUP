package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"FINANZAS/config"
	"FINANZAS/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func ObtenerCategorias(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_categoria, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
		FROM categoria
		ORDER BY id_categoria
	`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al consultar categoria")
		return
	}
	defer rows.Close()

	categorias := []models.Categoria{}

	for rows.Next() {
		categoria, err := scanCategoria(rows.Scan)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "error al leer categoria")
			return
		}
		categorias = append(categorias, categoria)
	}

	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "error al recorrer categoria")
		return
	}

	writeJSON(w, http.StatusOK, categorias)
}

func ObtenerCategoriaPorID(w http.ResponseWriter, r *http.Request) {
	idCategoria, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	row := config.DB.QueryRow(`
		SELECT id_categoria, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
		FROM categoria
		WHERE id_categoria = $1
	`, idCategoria)

	categoria, err := scanCategoria(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "categoria no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al consultar categoria")
		return
	}

	writeJSON(w, http.StatusOK, categoria)
}

func CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		INSERT INTO categoria (nombre, descripcion, activo)
		VALUES ($1, $2, $3)
		RETURNING id_categoria, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
	`, categoria.Nombre, categoria.Descripcion, categoria.Activo)

	categoriaCreada, err := scanCategoria(row.Scan)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al crear categoria")
		return
	}

	writeJSON(w, http.StatusCreated, categoriaCreada)
}

func ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	idCategoria, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	var categoria models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	row := config.DB.QueryRow(`
		UPDATE categoria
		SET nombre = $1,
			descripcion = $2,
			activo = $3
		WHERE id_categoria = $4
		RETURNING id_categoria, nombre, descripcion, activo, fecha_creacion, fecha_modificacion
	`, categoria.Nombre, categoria.Descripcion, categoria.Activo, idCategoria)

	categoriaActualizada, err := scanCategoria(row.Scan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "categoria no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al actualizar categoria")
		return
	}

	writeJSON(w, http.StatusOK, categoriaActualizada)
}

func EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	idCategoria, err := getIDFromRequest(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "id invalido")
		return
	}

	result, err := config.DB.Exec(`DELETE FROM categoria WHERE id_categoria = $1`, idCategoria)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			writeError(w, http.StatusConflict, "no se puede eliminar categoria porque tiene registros relacionados")
			return
		}
		writeError(w, http.StatusInternalServerError, "error al eliminar categoria")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error al validar eliminacion de categoria")
		return
	}

	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "categoria no encontrada")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "categoria eliminada correctamente",
	})
}

func scanCategoria(scan func(dest ...any) error) (models.Categoria, error) {
	var categoria models.Categoria
	var descripcion sql.NullString

	err := scan(
		&categoria.IdCategoria,
		&categoria.Nombre,
		&descripcion,
		&categoria.Activo,
		&categoria.FechaCreacion,
		&categoria.FechaModificacion,
	)
	if err != nil {
		return models.Categoria{}, err
	}

	categoria.Descripcion = nullStringToPointer(descripcion)
	return categoria, nil
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
