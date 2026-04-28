package controllers

import (
	"encoding/json"
	"net/http"
)

// respondJSON es el helper global para respuestas JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}