package models

import "time"

type Leccion struct {
	IdLeccion          int        `json:"id_leccion"`
	IdModulo           int        `json:"id_modulo"`
	IdContenido        *int       `json:"id_contenido"`
	Titulo             string     `json:"titulo"`
	Descripcion        *string    `json:"descripcion"`
	DuracionMinutos    *int       `json:"duracion_minutos"`
	UrlVideo           *string    `json:"url_video"`
	NumeroLeccion      *int       `json:"numero_leccion"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
