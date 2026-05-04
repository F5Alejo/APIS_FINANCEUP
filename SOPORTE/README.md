# Módulo de Soporte y PQR (SOPORTE)

Este microservicio se encarga de la atención al usuario, gestión de Peticiones, Quejas y Reclamos (PQR), y auditoría de actividades en FinanceUp.

## Entidades y Rutas Principales

El servicio abarca la gestión de:
- **PQR**: Creación, seguimiento y gestión de casos de soporte (`pqr_routes`, `estado_pqr_routes`).
- **Adjuntos**: Manejo de archivos e imágenes anexas a las solicitudes (`adjunto_routes`).
- **Registro de Actividad**: Trazabilidad y logs de acciones de los usuarios en la plataforma (`registro_actividad_routes`).

## Estructura del Proyecto

- `/config`: Configuración de base de datos y variables de entorno.
- `/controller`: Controladores de los endpoints.
- `/models`: Modelos de datos de soporte.
- `/routes`: Rutas expuestas para el módulo.

## Tecnologías
- Golang
- Gorilla Mux (Enrutador)
- PostgreSQL (Base de datos compartida, esquema propio)

[← Volver al README principal](../README.md)
