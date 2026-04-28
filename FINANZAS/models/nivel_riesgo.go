package models

import "time"

type NivelRiesgo struct {
	IdNivelRiesgo     int       `json:"id_nivel_riesgo"`
	Nombre            string    `json:"nombre"`
	Activo            bool      `json:"activo"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
}
