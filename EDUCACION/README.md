# Módulo Educativo Financiero (EDUCACION)

Este microservicio gestiona el contenido educativo financiero de FinanceUp, permitiendo a los usuarios acceder a módulos, lecciones y realizar un seguimiento de su progreso.

## Entidades y Rutas Principales

El servicio administra las siguientes entidades a través de su API:
- **Contenido**: Material y contenido detallado de estudio.
- **Lección**: Unidades de aprendizaje individuales.
- **Módulo Educativo**: Agrupaciones de lecciones por temas financieros.
- **Progreso Educativo**: Avance general del usuario en el programa.
- **Progreso Lección**: Estado y finalización de lecciones específicas.

## Estructura del Proyecto

- `/config`: Configuración de base de datos y entorno.
- `/controllers`: Lógica de manejo de peticiones HTTP.
- `/models`: Definición de estructuras de datos y modelos.
- `/routes`: Definición de los endpoints de la API.

## Tecnologías
- Golang
- Gorilla Mux (Enrutador)
- PostgreSQL (Base de datos compartida, esquema propio)

[← Volver al README principal](../README.md)
