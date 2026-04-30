package models

type TransaccionComision struct {
	ID_Transaccion     int     `json:"id_transaccion"`
	ID_Credito         int     `json:"id_credito"`
	ID_Banco           int     `json:"id_banco"`
	MontoComision      float64 `json:"monto_comision"`
	PorcentajeAplicado float64 `json:"porcentaje_aplicado"`
	FechaTransaccion   string  `json:"fecha_transaccion"`
	Estado             string  `json:"estado"`
	ReferenciaPago     string  `json:"referencia_pago"`
	Activo             bool    `json:"activo"`
}