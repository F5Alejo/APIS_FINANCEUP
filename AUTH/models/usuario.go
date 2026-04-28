package models

type Usuario struct {
	IDUsuario       int    `json:"id_usuario"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Cedula          string `json:"cedula"`
	Ciudad          string `json:"ciudad"`
	Estado          string `json:"estado"`
	IDTipoDocumento int    `json:"id_tipo_documento"`
	Activo          bool   `json:"activo"`
}
