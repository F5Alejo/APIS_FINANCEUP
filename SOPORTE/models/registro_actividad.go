package models

import "time"

// RegistroActividad representa la tabla soporte.registro_actividad
type RegistroActividad struct {
	ID               int        `json:"id_actividad"`
	IDUsuario        *int       `json:"id_usuario"`
	TipoActividad    *string    `json:"tipo_actividad"`
	Descripcion      *string    `json:"descripcion"`
	EntidadAfectada  *string    `json:"entidad_afectada"`
	FechaActividad   time.Time  `json:"fecha_actividad"`
	FechaCreacion    time.Time  `json:"fecha_creacion"`
}
