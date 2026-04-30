# FinanceUp — APIs Backend

Plataforma de gestión financiera personal que permite a los usuarios administrar sus finanzas, inversiones y metas económicas, acceder a contenido educativo financiero y recibir soporte a través de un sistema de PQR.

El backend está compuesto por **4 microservicios independientes**, cada uno con su propio esquema en la base de datos PostgreSQL `FinanceUp`.

---

## Arquitectura General

```
APIS_FINANCEUP/
├── AUTH/          → Autenticación y autorización             (puerto: en desarrollo)
├── EDUCACION/     → Módulo educativo financiero               (puerto: 8080)
├── FINANZAS/      → Gestión financiera, metas e inversiones   (puerto: 8081)
├── NEGOCIO/       → Lógica de negocio                         (puerto: en desarrollo)
└── SOPORTE/       → PQR, adjuntos y registro de actividad     (puerto: 8083)
```

Todos los servicios comparten la misma base de datos PostgreSQL (`FinanceUp`) pero operan sobre **esquemas separados**, garantizando aislamiento de datos entre módulos.

---

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang 1.24+](https://go.dev/doc/install)
- [Gorilla Mux v1.8.1](https://github.com/gorilla/mux) — enrutador HTTP
- [lib/pq v1.12.3](https://github.com/lib/pq) — driver PostgreSQL para Go
- [PostgreSQL](https://www.postgresql.org/) — base de datos relacional (esquemas por módulo)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

---

### Variables de Entorno

Cada microservicio puede configurarse mediante variables de entorno. Los valores por defecto están definidos en cada `config/db.go`.

```
# Base de datos (compartida por todos los servicios)
FINANCEUP_PGHOST=[dirección del servidor PostgreSQL]        # default: localhost
FINANCEUP_PGPORT=[puerto de conexión]                       # default: 5432
FINANCEUP_PGUSER=[usuario con acceso a la base de datos]    # default: postgres
FINANCEUP_PGPASS=[contraseña del usuario]
FINANCEUP_PGDB=[nombre de la base de datos]                 # default: FinanceUp

# Puertos HTTP por servicio
EDUCACION_HTTP_PORT=[puerto del servicio educación]         # default: 8080
FINANZAS_HTTP_PORT=[puerto del servicio finanzas]           # default: 8081
SOPORTE_HTTP_PORT=[puerto del servicio soporte]             # default: 8083
```

---

## Ejecución de los Servicios

### Ejecución Local

```bash
# 1. Clonar el repositorio
git clone -b develop https://github.com/F5Alejo/APIS_FINANCEUP

# 2. Moverse a la carpeta del repositorio
cd APIS_FINANCEUP
