package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"FINANZAS/config"

	"github.com/gorilla/mux"
)

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

func nullStringToPointer(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	text := value.String
	return &text
}

func nullInt32ToPointer(value sql.NullInt32) *int {
	if !value.Valid {
		return nil
	}

	number := int(value.Int32)
	return &number
}

func nullFloat64ToPointer(value sql.NullFloat64) *float64 {
	if !value.Valid {
		return nil
	}

	number := value.Float64
	return &number
}

func nullBoolToPointer(value sql.NullBool) *bool {
	if !value.Valid {
		return nil
	}

	result := value.Bool
	return &result
}

func nullTimeToPointer(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	fecha := value.Time
	return &fecha
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
