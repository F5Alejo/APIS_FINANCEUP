# Módulo de Gestión Financiera (FINANZAS)

Este microservicio maneja la lógica central de las finanzas personales en FinanceUp, incluyendo la gestión de ingresos, egresos, metas de ahorro e inversiones.

## Entidades y Rutas Principales

El servicio expone endpoints para gestionar:
- **Finanzas Generales**: Gestión de categorías y el perfil financiero.
- **Ingresos y Egresos**: Registro de movimientos (`movimiento_ingreso_egreso`, `tipo_ingreso`).
- **Metas Financieras**: Creación y seguimiento de objetivos (`meta`, `editar_meta`, `movimiento_meta`, `tipo_ingreso_meta`).
- **Inversiones**: Portafolio de inversiones (`inversion`, `movimiento_inversion`, `tipo_inversion`, `tipo_ingreso_inversion`, `nivel_riesgo`).

## Estructura del Proyecto

- `/config`: Configuración de base de datos y entorno.
- `/controllers`: Controladores para el procesamiento de solicitudes.
- `/models`: Representación de datos financieros.
- `/routes`: Definición y agrupación de rutas HTTP.

## Tecnologías
- Golang
- Gorilla Mux (Enrutador)
- PostgreSQL (Base de datos compartida, esquema propio)

[← Volver al README principal](../README.md)
