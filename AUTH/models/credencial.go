package models

type Credencial struct {
	IDCredencial    int    `json:"id_credencial"`
	IDUsuario       int    `json:"id_usuario"`
	ContrasenaHash  string `json:"contrasena_hash"`
	IntentosFallidos int   `json:"intentos_fallidos"`
	RequiereCambio  bool   `json:"requiere_cambio"`
	Activo          bool   `json:"activo"`
}
