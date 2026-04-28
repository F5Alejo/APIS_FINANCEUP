package models

import "time"

// EstadoPqr representa la tabla soporte.estado_pqr
type EstadoPqr struct {
	ID                 int        `json:"id_estado"`
	Nombre             string     `json:"nombre"`
	Descripcion        *string    `json:"descripcion"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
