package models

import "time"

type EditarMeta struct {
	IdEditarMeta      int        `json:"id_editar_meta"`
	Nombre            *string    `json:"nombre"`
	MontoActual       *float64   `json:"monto_actual"`
	MontoObjetivo     float64    `json:"monto_objetivo"`
	AhorroMensual     *float64   `json:"ahorro_mensual"`
	FechaObjetivo     *time.Time `json:"fecha_objetivo"`
	Descripcion       *string    `json:"descripcion"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
