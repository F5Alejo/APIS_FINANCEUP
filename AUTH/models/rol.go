package models

type Rol struct {
	IDRol       int    `json:"id_rol"`
	NombreRol   string `json:"nombre_rol"`
	Descripcion string `json:"descripcion"`
	Activo      bool   `json:"activo"`
}
