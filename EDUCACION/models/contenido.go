package models

import "time"

type Contenido struct {
	IdContenido       int        `json:"id_contenido"`
	Titulo            string     `json:"titulo"`
	Descripcion       *string    `json:"descripcion"`
	DuracionMinutos   *int       `json:"duracion_minutos"`
	UrlVideo          *string    `json:"url_video"`
	Activo            bool       `json:"activo"`
	FechaCreacion     time.Time  `json:"fecha_creacion"`
	FechaModificacion time.Time  `json:"fecha_modificacion"`
}
