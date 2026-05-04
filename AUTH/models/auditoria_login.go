package models

type AuditoriaLogin struct {
	IDAuditoria  int    `json:"id_auditoria"`
	IDUsuario    int    `json:"id_usuario"`
	EstadoEvento string `json:"estado_evento"`
	Navegador    string `json:"navegador"`
	FechaEvento  string `json:"fecha_evento"`
}
