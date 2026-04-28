package models

type ConversacionUsuarioAsesor struct {
	ID_Conversacion int    `json:"id_conversacion"`
	ID_Lead         int    `json:"id_lead"`
	ID_Usuario      int    `json:"id_usuario"`
	ID_Asesor       int    `json:"id_asesor"`
	TipoContacto    string `json:"tipo_contacto"`
	Asunto          string `json:"asunto"`
	Contenido       string `json:"contenido"`
	FechaMensaje    string `json:"fecha_mensaje"`
	Activo          bool   `json:"activo"`
}