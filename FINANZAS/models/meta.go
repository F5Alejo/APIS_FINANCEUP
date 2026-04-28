package models

import "time"

type Meta struct {
	IdMeta             int        `json:"id_meta"`
	IdUsuario          int        `json:"id_usuario"`
	IdEditarMeta       int        `json:"id_editar_meta"`
	IdMovimientoDinero *int       `json:"id_movimiento_dinero"`
	Nombre             *string    `json:"nombre"`
	Descripcion        *string    `json:"descripcion"`
	MontoObjetivo      *float64   `json:"monto_objetivo"`
	MontoActual        *float64   `json:"monto_actual"`
	FechaLimite        *time.Time `json:"fecha_limite"`
	Color              *string    `json:"color"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
