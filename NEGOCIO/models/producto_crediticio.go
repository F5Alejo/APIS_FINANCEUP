package models

type ProductoCrediticio struct {
	ID             int     `json:"id_producto"`
	IDBanco        int     `json:"id_banco"`
	NombreProducto string  `json:"nombre_producto"`
	Descripcion    string  `json:"descripcion"`
	MontoMinimo    float64 `json:"monto_minimo"`
	MontoMaximo    float64 `json:"monto_maximo"`
	TasaMinima     float64 `json:"tasa_minima"`
	TasaMaxima     float64 `json:"tasa_maxima"`
	PlazoMinimo    int     `json:"plazo_minimo"`
	PlazoMaximo    int     `json:"plazo_maximo"`
	Requisitos     string  `json:"requisitos"`
	Activo         bool    `json:"activo"`
}