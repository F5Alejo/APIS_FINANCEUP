package models

import "time"

type TipoInversion struct {
	IdTipoInversion   int       `json:"id_tipo_inversion"`
	Nombre            string    `json:"nombre"`
	Descripcion       *string   `json:"descripcion"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
