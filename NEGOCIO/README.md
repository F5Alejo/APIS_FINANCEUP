# Módulo de Lógica de Negocio (NEGOCIO)

Este microservicio gestiona las interacciones comerciales, leads, créditos y alianzas con entidades bancarias dentro de FinanceUp.

## Entidades y Rutas Principales

El servicio soporta las siguientes operaciones a través de su API:
- **Bancos y Productos**: Gestión de entidades bancarias y productos crediticios (`banco_routes`, `producto_crediticio`).
- **Asesores**: Gestión de asesores bancarios y contacto directo (`asesor_bancario`, `contacto_asesor`, `conversacion_usuario_asesor`).
- **Leads y Créditos**: Seguimiento de prospectos y créditos desembolsados (`leads`, `credito_desembolsado`).
- **Comisiones**: Registro de transacciones y comisiones generadas (`transaccion_comision`).

## Estructura del Proyecto

- `/config`: Configuración de base de datos.
- `/controllers`: Lógica para el manejo de endpoints de negocio.
- `/models`: Estructuras de datos relacionadas a ventas y leads.
- `/routes`: Registro de rutas de los servicios.

## Tecnologías
- Golang
- Gorilla Mux (Enrutador)
- PostgreSQL (Base de datos compartida, esquema propio)

[← Volver al README principal](../README.md)
