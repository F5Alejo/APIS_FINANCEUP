package models

type UsuarioRol struct {
	IDUsuarioRol    int    `json:"id_usuario_rol"`
	IDUsuario       int    `json:"id_usuario"`
	IDRol           int    `json:"id_rol"`
	FechaAsignacion string `json:"fecha_asignacion"`
	Activo          bool   `json:"activo"`
}
