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

## Licencia

Proyecto educativo - Escuela Simón Bolívar 4-084
