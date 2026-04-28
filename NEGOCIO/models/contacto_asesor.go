package models

type ContactoAsesor struct {
	ID_Contacto       int    `json:"id_contacto"`
	ID_Asesor         int    `json:"id_asesor"`
	Whatsapp          string `json:"whatsapp"`
	Email             string `json:"email"`
	Telefono          string `json:"telefono"`
	DisponibleDesde   string `json:"disponible_desde"`
	DisponibleHasta   string `json:"disponible_hasta"`
	DiasDisponibles   string `json:"dias_disponibles"`
	Activo            bool   `json:"activo"`
}