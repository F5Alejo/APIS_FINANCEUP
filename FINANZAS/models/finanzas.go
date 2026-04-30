package models

import "time"

type Finanzas struct {
	IdFinanzas         int        `json:"id_finanzas"`
	IdUsuario          int        `json:"id_usuario"`
	IdMovimientoDinero *int       `json:"id_movimiento_dinero"`
	IdCategoria        *int       `json:"id_categoria"`
	MontoPresupuesto   *float64   `json:"monto_presupuesto"`
	Gasto              *float64   `json:"gasto"`
	Disponible         *float64   `json:"disponible"`
	Fecha              *time.Time `json:"fecha"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
