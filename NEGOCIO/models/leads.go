package models

type Lead struct {
	ID_Lead          int     `json:"id_lead"`
	ID_Usuario       int     `json:"id_usuario"`
	ID_Producto      int     `json:"id_producto"`
	ID_Asesor        int     `json:"id_asesor"`
	TipoCredito      string  `json:"tipo_credito"`
	MontoInteres     float64 `json:"monto_interes"`
	PlazoInteres     int     `json:"plazo_interes"`
	EstadoLead       string  `json:"estado_lead"`
	FechaGeneracion  string  `json:"fecha_generacion"`
	FechaContacto    string  `json:"fecha_contacto"`
	Observaciones    string  `json:"observaciones"`
	Activo           bool    `json:"activo"`
}