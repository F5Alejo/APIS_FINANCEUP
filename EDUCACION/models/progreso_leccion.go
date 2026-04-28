package models

import "time"

type ProgresoLeccion struct {
	IdProgresoLeccion   int        `json:"id_progreso_leccion"`
	IdUsuario           int        `json:"id_usuario"`
	IdLeccion           int        `json:"id_leccion"`
	Completado          *bool      `json:"completado"`
	FechaInicio         *time.Time `json:"fecha_inicio"`
	FechaCompletado     *time.Time `json:"fecha_completado"`
	Activo              bool       `json:"activo"`
	FechaCreacion       time.Time  `json:"fecha_creacion"`
	FechaModificacion   time.Time  `json:"fecha_modificacion"`
}
