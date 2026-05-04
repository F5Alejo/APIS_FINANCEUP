# Módulo de Autenticación y Autorización (AUTH)

Este microservicio se encarga de gestionar la autenticación y autorización de los usuarios en la plataforma FinanceUp.

## Entidades y Rutas Principales

El servicio administra las siguientes entidades a través de su API:
- **Usuarios**: Gestión del registro y perfil básico de usuarios (`usuario`).
- **Credenciales**: Administración de contraseñas y accesos (`credencial`).
- **Roles**: Definición de niveles de acceso en la plataforma (`rol`, `usuario_rol`).
- **Tipos de Documento**: Catálogo de tipos de identificación (`tipo_documento`).
- **Auditoría de Login**: Registro y monitoreo de inicios de sesión (`auditoria_login`).

## Estructura del Proyecto

- `/config`: Configuración de base de datos y entorno.
- `/controller`: Lógica de manejo de peticiones HTTP.
- `/models`: Definición de estructuras de datos y modelos.
- `/router`: Definición de los endpoints de la API.

## Tecnologías
- Golang
- Gorilla Mux (Enrutador)
- PostgreSQL (Base de datos compartida, esquema propio)

[← Volver al README principal](../README.md)
