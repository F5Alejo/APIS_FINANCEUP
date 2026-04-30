package models

import "time"

type ModuloEducativo struct {
	IdModulo          int        `json:"id_modulo"`
	Titulo            string     `json:"titulo"`
	Descripcion       *string    `json:"descripcion"`
	Contenido         *string    `json:"contenido"`
	Nivel             string     `json:"nivel"`
	UrlThumbnail      *string    `json:"url_thumbnail"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
