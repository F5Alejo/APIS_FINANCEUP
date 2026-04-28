package models

import "time"

type MovimientoMeta struct {
	IdMovimientoDinero int       `json:"id_movimiento_dinero"`
	Nombre             string    `json:"nombre"`
	Monto              float64   `json:"monto"`
	EsIngreso          bool      `json:"es_ingreso"`
	Activo             bool      `json:"activo"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaModificacion  time.Time `json:"fecha_modificacion"`
}
