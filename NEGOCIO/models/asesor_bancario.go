package models

// AsesorBancario representa la tabla negocio.asesor_bancario
type AsesorBancario struct {
	IDAsesor     int    `json:"id_asesor"`
	IDBanco      int    `json:"id_banco"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Email        string `json:"email"`
	Telefono     string `json:"telefono"`
	Especialidad string `json:"especialidad"`
	Activo       bool   `json:"activo"`
}