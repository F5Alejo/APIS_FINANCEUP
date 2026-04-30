package models

import "time"

type Categoria struct {
	IdCategoria       int       `json:"id_categoria"`
	Nombre            string    `json:"nombre"`
	Descripcion       *string   `json:"descripcion"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
