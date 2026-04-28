package controllers

import (
	"NEGOCIO/config"
	"NEGOCIO/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllConversaciones(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_conversacion, id_lead, id_usuario, id_asesor,
	           tipo_contacto, asunto, contenido, fecha_mensaje, activo
	           FROM negocio.conversacion_usuario_asesor WHERE 1=1`

	tipo := r.URL.Query().Get("tipo_contacto")
	if tipo != "" {
		query += " AND tipo_contacto ILIKE '%" + tipo + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.ConversacionUsuarioAsesor
	for rows.Next() {
		var c models.ConversacionUsuarioAsesor
		rows.Scan(&c.ID_Conversacion, &c.ID_Lead, &c.ID_Usuario, &c.ID_Asesor,
			&c.TipoContacto, &c.Asunto, &c.Contenido, &c.FechaMensaje, &c.Activo)
		list = append(list, c)
	}
	respondJSON(w, 200, list)
}

func GetConversacionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.ConversacionUsuarioAsesor

	err := config.DB.QueryRow(
		`SELECT id_conversacion, id_lead, id_usuario, id_asesor,
		  tipo_contacto, asunto, contenido, fecha_mensaje, activo
		  FROM negocio.conversacion_usuario_asesor WHERE id_conversacion = $1`, id,
	).Scan(&c.ID_Conversacion, &c.ID_Lead, &c.ID_Usuario, &c.ID_Asesor,
		&c.TipoContacto, &c.Asunto, &c.Contenido, &c.FechaMensaje, &c.Activo)

	if err == sql.ErrNoRows {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, c)
}

func CreateConversacion(w http.ResponseWriter, r *http.Request) {
	var c models.ConversacionUsuarioAsesor
	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		`INSERT INTO negocio.conversacion_usuario_asesor
		 (id_lead, id_usuario, id_asesor, tipo_contacto, asunto, contenido, activo)
		 VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_conversacion`,
		c.ID_Lead, c.ID_Usuario, c.ID_Asesor,
		c.TipoContacto, c.Asunto, c.Contenido, c.Activo,
	).Scan(&c.ID_Conversacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, c)
}

func UpdateConversacion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var c models.ConversacionUsuarioAsesor
	json.NewDecoder(r.Body).Decode(&c)

	_, err := config.DB.Exec(
		`UPDATE negocio.conversacion_usuario_asesor
		 SET id_lead=$1, id_usuario=$2, id_asesor=$3,
		     tipo_contacto=$4, asunto=$5, contenido=$6, activo=$7
		 WHERE id_conversacion=$8`,
		c.ID_Lead, c.ID_Usuario, c.ID_Asesor,
		c.TipoContacto, c.Asunto, c.Contenido, c.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Actualizado"})
}

func DeleteConversacion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := config.DB.Exec("DELETE FROM negocio.conversacion_usuario_asesor WHERE id_conversacion=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"message": "Dato Eliminado"})
}