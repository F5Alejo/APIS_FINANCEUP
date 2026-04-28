package models

import "time"

// Pqr representa la tabla soporte.pqr
type Pqr struct {
	ID                 int       `json:"id_pqr"`
	IDUsuario          int       `json:"id_usuario"`
	Descripcion        string    `json:"descripcion"`
	IDEstado           int       `json:"id_estado"`
	Activo             bool      `json:"activo"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaModificacion  time.Time `json:"fecha_modificacion"`
}
