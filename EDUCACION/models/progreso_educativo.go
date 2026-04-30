package models

import "time"

type ProgresoEducativo struct {
	IdProgreso            int        `json:"id_progreso"`
	IdUsuario             int        `json:"id_usuario"`
	IdModulo              int        `json:"id_modulo"`
	PorcentajeCompletado  *int       `json:"porcentaje_completado"`
	FechaInicio           *time.Time `json:"fecha_inicio"`
	FechaCompletado       *time.Time `json:"fecha_completado"`
	Calificacion          *int       `json:"calificacion"`
	Activo                bool       `json:"activo"`
	FechaCreacion         time.Time  `json:"fecha_creacion"`
	FechaModificacion     time.Time  `json:"fecha_modificacion"`
}
