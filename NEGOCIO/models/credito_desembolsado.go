package models

type CreditoDesembolsado struct {
	ID_Credito        int     `json:"id_credito"`
	ID_Lead           int     `json:"id_lead"`
	ID_Usuario        int     `json:"id_usuario"`
	ID_Producto       int     `json:"id_producto"`
	ID_Banco          int     `json:"id_banco"`
	NumeroCredito     string  `json:"numero_credito"`
	MontoAprobado     float64 `json:"monto_aprobado"`
	TasaInterestFinal float64 `json:"tasa_interes_final"`
	PlazoMeses        int     `json:"plazo_meses"`
	FechaAprobacion   string  `json:"fecha_aprobacion"`
	FechaDesembolso   string  `json:"fecha_desembolso"`
	EstadoCredito     string  `json:"estado_credito"`
	SaldoActual       float64 `json:"saldo_actual"`
	Activo            bool    `json:"activo"`
}