package models

import "time"

type Inversion struct {
	IdInversion        int        `json:"id_inversion"`
	IdUsuario          int        `json:"id_usuario"`
	IdTipoInversion    int        `json:"id_tipo_inversion"`
	IdNivelRiesgo      int        `json:"id_nivel_riesgo"`
	IdMovimientoDinero *int       `json:"id_movimiento_dinero"`
	Nombre             *string    `json:"nombre"`
	Monto              *float64   `json:"monto"`
	Rentabilidad       *float64   `json:"rentabilidad"`
	FechaInicio        *time.Time `json:"fecha_inicio"`
	FechaFin           *time.Time `json:"fecha_fin"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
