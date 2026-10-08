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
| Frontend | Angular 19 |
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

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/api/students` | Listar estudiantes |
| POST | `/api/students` | Crear estudiante |
| GET | `/api/students/{id}` | Obtener estudiante |
| PUT | `/api/students/{id}` | Actualizar estudiante |
| DELETE | `/api/students/{id}` | Eliminar estudiante |
| GET | `/api/groups` | Listar grupos |
| POST | `/api/groups` | Crear grupo |
| GET | `/api/groups/{id}` | Obtener grupo |
| PUT | `/api/groups/{id}` | Actualizar grupo |
| DELETE | `/api/groups/{id}` | Eliminar grupo |
| GET | `/api/invitations` | Listar invitaciones |
| POST | `/api/invitations` | Crear invitación |
| PUT | `/api/invitations/{groupId}` | Actualizar invitación |

## Licencia

Proyecto educativo - Escuela Simón Bolívar 4-084
