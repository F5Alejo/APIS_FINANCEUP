package models

// Banco representa la tabla negocio.banco
type Banco struct {
	IDBanco            int     `json:"id_banco"`
	NombreBanco        string  `json:"nombre_banco"`
	Ciudad             string  `json:"ciudad"`
	Contacto           string  `json:"contacto"`
	Telefono           string  `json:"telefono"`
	Email              string  `json:"email"`
	ComisionPorcentaje float64 `json:"comision_porcentaje"`
	UrlLogo            string  `json:"url_logo"`
	Descripcion        string  `json:"descripcion"`
	SitioWeb           string  `json:"sitio_web"`
	Estado             string  `json:"estado"`
	Activo             bool    `json:"activo"`
}