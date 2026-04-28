package models

import "time"

// Adjunto representa la tabla soporte.adjunto
type Adjunto struct {
	ID                 int        `json:"id_adjunto"`
	IDPqr              int        `json:"id_pqr"`
	NombreArchivo      string     `json:"nombre_archivo"`
	RutaArchivo        string     `json:"ruta_archivo"`
	TipoMime           *string    `json:"tipo_mime"`
	TamanoBytes        *int       `json:"tamano_bytes"`
	FechaCarga         time.Time  `json:"fecha_carga"`
	Activo             bool       `json:"activo"`
	FechaCreacion      time.Time  `json:"fecha_creacion"`
	FechaModificacion  time.Time  `json:"fecha_modificacion"`
}
