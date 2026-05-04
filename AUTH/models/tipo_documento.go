package models

type TipoDocumento struct {
	TipoDocumento int    `json:"tipo_documento"`
	Nombre          string `json:"nombre"`
	Codigo          string `json:"codigo"`
	Activo          bool   `json:"activo"`
}
