package models

type TipoDocumento struct {
	IDTipoDocumento int    `json:"id_tipo_documento"`
	Nombre          string `json:"nombre"`
	Codigo          string `json:"codigo"`
	Activo          bool   `json:"activo"`
}
