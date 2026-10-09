# Inscripciones Simón Bolívar 4-084

Sistema web para los ingresantes al ciclo lectivo escolar de la escuela secundaria Simón Bolívar 4-084.

## Descripción

En este sistema los ingresantes podrán:

- **Confirmar sus bancos**: Verificar su inscripción y reservar su lugar
- **Formar grupos**: Crear o unirse a grupos para comenzar su cursado

## Estructura del Proyecto

```
inscripciones-sb-web/
├── frontend/          # Aplicación Angular
│   ├── src/
│   │   ├── app/
│   │   │   ├── features/    # Componentes por funcionalidad
│   │   │   ├── models/      # Modelos de datos
│   │   │   ├── services/    # Servicios HTTP
│   │   │   └── shared/      # Componentes compartidos
│   └── ...
└── backend/           # API Go + MongoDB
    ├── config/        # Configuración de base de datos
    ├── models/        # Modelos de datos
    ├── repository/    # Operaciones CRUD
    ├── handlers/      # Manejadores HTTP
    └── router/        # Rutas de la API
```

## Tecnologías

| Capa | Tecnología |
|------|------------|
| Frontend | Angular 22 |
| Backend | Go 1.21+ |
| Base de datos | MongoDB Atlas |

## Configuración

### Frontend

```bash
cd frontend
npm install
ng serve
```

### Backend

```bash
cd backend
export MONGODB_URI="mongodb+srv://usuario:password@cluster0.xxxxx.mongodb.net/"
go run main.go
```

## Endpoints de la API

### Públicas (sin autenticación)

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/api/login` | Iniciar sesión del administrador |
| GET | `/api/students` | Listar estudiantes |
| GET | `/api/students/{id}` | Obtener estudiante |
| POST | `/api/students/{id}/confirm` | Confirmar banco (flujo del ingresante) |
| GET | `/api/groups` | Listar grupos |
| GET | `/api/groups/{id}` | Obtener grupo |
| GET | `/api/invitations` | Listar invitaciones |

### Protegidas (requieren `Authorization: Bearer <token>`)

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/api/admin/students` | Listar estudiantes (dashboard) |
| POST | `/api/admin/students` | Crear estudiante |
| PUT | `/api/admin/students/{id}` | Actualizar estudiante |
| DELETE | `/api/admin/students/{id}` | Eliminar estudiante |
| POST | `/api/admin/students/import` | Importar alumnos desde Excel |
| POST | `/api/admin/groups` | Crear grupo |
| PUT | `/api/admin/groups/{id}` | Actualizar grupo |
| DELETE | `/api/admin/groups/{id}` | Eliminar grupo |
| DELETE | `/api/admin/groups/{id}/members/{memberId}` | Quitar integrante |

> Toda escritura requiere token de administrador. Solo la lectura de listados y la
> confirmación de banco son públicas, porque las consulta la comunidad sin cuenta.

## Variables de entorno

### Backend (`backend/.env`, ver `backend/.env.example`)

| Variable | Descripción |
|----------|-------------|
| `MONGODB_URI` | URI de MongoDB Atlas |
| `ADMIN_USERNAME` | Usuario único del dashboard |
| `ADMIN_PASSWORD` | Hash **bcrypt** de la contraseña, entre comillas simples |
| `JWT_SECRET` | Clave para firmar los JWT |
| `ALLOWED_ORIGINS` | Orígenes del frontend separados por coma. Vacío = cualquier origen (solo dev) |
| `PORT` | Puerto de escucha. Render lo inyecta solo |

Para generar el hash bcrypt de la contraseña:

```bash
htpasswd -bnBC 10 "" TU_CONTRASENA | tr -d ':\n'
```

### Frontend

| Variable | Cuándo | Descripción |
|----------|--------|-------------|
| `API_URL` | Solo en build de Render | URL base de la API. El script `scripts/set-api-url.js` la inyecta en `environment.prod.ts` |

En desarrollo no hace falta: `ng serve` usa `environment.ts`, que ya apunta a
`http://localhost:8080/api`.

Para compilar a mano contra otro destino:

```bash
cd frontend
API_URL=https://mi-api.onrender.com npm run build:render
```

## Despliegue en Render

El repositorio incluye un [`render.yaml`](./render.yaml) que define los dos servicios.

1. Subí el código a GitHub (el `.env` ya está en `.gitignore`).
2. En Render: **New → Blueprint**, apuntando al repositorio.
3. Render te va a pedir los valores `sync: false`:
   - `MONGODB_URI` — la URI de Atlas
   - `ADMIN_USERNAME` — el usuario del dashboard
   - `ADMIN_PASSWORD` — el hash bcrypt entre comillas simples
4. `JWT_SECRET` se genera solo.
5. `API_URL` y `ALLOWED_ORIGINS` se completan automáticamente con los hostnames
   públicos de cada servicio.

### Qué queda como está

| Servicio | Plan | Build | Start | Publicación |
|----------|------|-------|-------|-------------|
| `inscripciones-sb-api` | free | `go build -o api .` | `./api` | — |
| `inscripciones-sb` | free (static) | `npm ci && npm run build:render` | — | `dist/sb-web/browser` |

El sitio estático declara una regla `rewrite /* → /index.html` para que las rutas
de Angular (`/fechas`, `/grupos`, `/admin`) funcionen al entrar directo por URL.

> **Nota sobre el plan gratuito**: los web services de Render se suspenden tras
> ~15 minutos de inactividad y tardan ~1 minuto en volver. La primera visita
> después de un rato pausado va a ser lenta. Si molesta en producción, conviene
> pasar el servicio de la API al plan Hobby.

## Licencia

Proyecto educativo - Escuela Simón Bolívar 4-084
