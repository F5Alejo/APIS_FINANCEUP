
# APIS_FINANCEUP — negocio

API REST en Go para la gestión del ecosistema bancario y comercial de la plataforma **FinanceUp**. Administra bancos, asesores, productos crediticios, leads, conversaciones, créditos desembolsados y comisiones sobre el esquema `negocio` de PostgreSQL.

---

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

* [Go 1.26.2](https://go.dev/)
* [Gorilla Mux v1.8.1](https://github.com/gorilla/mux)
* [lib/pq v1.12.3](https://github.com/lib/pq) — driver PostgreSQL para Go
* [PostgreSQL](https://www.postgresql.org/) — base de datos `FINANCEUP`, esquema `negocio`

---

### Variables de Entorno

Configurar las siguientes variables antes de ejecutar el proyecto (ver `config/db.go`):

```
NEGOCIO_PGHOST=[dirección del servidor PostgreSQL]
NEGOCIO_PGPORT=[puerto de conexión, por defecto 5432]
NEGOCIO_PGUSER=[usuario con acceso a la base de datos]
NEGOCIO_PGPASS=[contraseña del usuario]
NEGOCIO_PGDB=[nombre de la base de datos, por defecto FINANCEUP]
NEGOCIO_PGSCHEMA=[esquema de tablas, por defecto negocio]
NEGOCIO_HTTP_PORT=[puerto de ejecución, por defecto 8085]
```

---

### Ejecución del Proyecto

```bash
# 1. Clonar el repositorio
git clone https://github.com/F5Alejo/APIS_FINANCEUP.git

# 2. Moverse a la carpeta del microservicio
cd APIS_FINANCEUP/NEGOCIO

# 3. Descargar dependencias
go mod tidy

# 4. Ejecutar el servidor
go run main.go
```

El servidor quedará escuchando en **`:8085`**.

---

### Ejecución Pruebas

Pruebas unitarias

```bash
# En proceso
```

---

## Modelo de Datos

El esquema `negocio` contiene las siguientes tablas:

| Tabla                        | Descripción                                              |
|------------------------------|----------------------------------------------------------|
| `banco`                      | Bancos aliados con contacto, comisión y estado           |
| `asesor_bancario`            | Asesores vinculados a un banco con especialidad          |
| `contacto_asesor`            | Canales de contacto y horarios de disponibilidad         |
| `producto_crediticio`        | Productos de crédito con rangos de monto, tasa y plazo   |
| `lead`                       | Solicitudes de crédito generadas por usuarios            |
| `conversacion_usuario_asesor`| Historial de mensajes entre usuario y asesor por lead    |
| `credito_desembolsado`       | Créditos aprobados con saldo, estado y fechas            |
| `transaccion_comision`       | Comisiones generadas por crédito desembolsado            |

---

## Endpoints Disponibles

Todos los endpoints responden en formato **JSON** y soportan `GET`, `POST`, `PUT`, `DELETE`.

| Recurso                  | Ruta base                        | Puerto |
|--------------------------|----------------------------------|--------|
| Bancos                   | `/bancos`                        | 8085   |
| Asesores bancarios       | `/negocio/asesores`              | 8085   |
| Productos crediticios    | `/negocio/productos`             | 8085   |
| Leads                    | `/negocio/leads`                 | 8085   |
| Conversaciones           | `/negocio/conversaciones`        | 8085   |
| Créditos desembolsados   | `/negocio/creditos`              | 8085   |
| Transacciones comisión   | `/negocio/transacciones`         | 8085   |

---

## Estado CI

| Develop | Release 0.0.1 | Master |
|---------|---------------|--------|
| En proceso | En proceso | En proceso |

---

## Licencia

This file is part of APIS_FINANCEUP.

APIS_FINANCEUP is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

APIS_FINANCEUP is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with APIS_FINANCEUP. If not, see <https://www.gnu.org/licenses/>.
