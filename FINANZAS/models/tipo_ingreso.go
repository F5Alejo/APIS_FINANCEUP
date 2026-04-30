package models

import "time"

type TipoIngreso struct {
	IdTipoIngreso        int       `json:"id_tipo_ingreso"`
	IdMovimientoDinero   *int      `json:"id_movimiento_dinero"`
	NombreMovimientoPago string    `json:"nombre_movimiento_pago"`
	Descripcion          *string   `json:"descripcion"`
	Activo               bool      `json:"activo"`
	FechaCreacion        time.Time `json:"fecha_creacion"`
	FechaModificacion    time.Time `json:"fecha_modificacion"`
}
