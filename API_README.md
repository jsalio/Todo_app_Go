# API de Gestión de Tareas (TODO)

Esta es una API RESTful para gestionar tareas, construida con Go y Gin, que utiliza Supabase como base de datos.

## Requisitos

- Go 1.16 o superior
- Cuenta de Supabase
- Variables de entorno configuradas en un archivo `.env`:
  ```
  SUPABASE_URL=tu_url_de_supabase
  SUPABASE_KEY=tu_clave_de_api_de_supabase
  PORT=8080 (opcional, por defecto es 8080)
  ```

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/todo-app-go.git
   cd todo-app-go
   ```

2. Instala las dependencias:
   ```bash
   go mod tidy
   ```

3. Copia el archivo `.env.example` a `.env` y configura tus credenciales de Supabase.

## Ejecución

Para iniciar el servidor de la API:

```bash
go run cmd/web/main.go
```

El servidor estará disponible en `http://localhost:8080` por defecto.

## Endpoints de la API

### Listar todas las tareas

```
GET /api/todos
```

**Respuesta exitosa (200 OK):**
```json
[
  {
    "id": 1,
    "task": "Comprar leche",
    "is_complete": false,
    "created_at": "2023-01-01T12:00:00Z"
  },
  ...
]
```

### Obtener una tarea por ID

```
GET /api/todos/:id
```

**Parámetros de la URL:**
- `id` - ID numérico de la tarea

**Respuesta exitosa (200 OK):**
```json
{
  "id": 1,
  "task": "Comprar leche",
  "is_complete": false,
  "created_at": "2023-01-01T12:00:00Z"
}
```

**Respuesta de error (404 Not Found):**
```json
{
  "error": "Tarea no encontrada"
}
```

### Crear una nueva tarea

```
POST /api/todos
```

**Cuerpo de la petición (JSON):**
```json
{
  "task": "Nueva tarea"
}
```

**Respuesta exitosa (201 Created):**
```json
{
  "id": 2,
  "task": "Nueva tarea",
  "is_complete": false,
  "created_at": "2023-01-02T10:30:00Z"
}
```

### Actualizar una tarea existente

```
PUT /api/todos/:id
```

**Parámetros de la URL:**
- `id` - ID numérico de la tarea a actualizar

**Cuerpo de la petición (JSON):**
```json
{
  "task": "Tarea actualizada",
  "is_complete": true
}
```

**Nota:** Ambos campos son opcionales. Solo se actualizarán los campos proporcionados.

**Respuesta exitosa (200 OK):**
```json
{
  "id": 1,
  "task": "Tarea actualizada",
  "is_complete": true,
  "created_at": "2023-01-01T12:00:00Z"
}
```

### Eliminar una tarea

```
DELETE /api/todos/:id
```

**Parámetros de la URL:**
- `id` - ID numérico de la tarea a eliminar

**Respuesta exitosa (204 No Content):**
```
(No hay contenido en la respuesta)
```

## Ejemplos de uso con cURL

### Listar todas las tareas
```bash
curl http://localhost:8080/api/todos
```

### Obtener una tarea específica
```bash
curl http://localhost:8080/api/todos/1
```

### Crear una nueva tarea
```bash
curl -X POST -H "Content-Type: application/json" -d '{"task":"Nueva tarea"}' http://localhost:8080/api/todos
```

### Actualizar una tarea
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"task":"Tarea actualizada","is_complete":true}' http://localhost:8080/api/todos/1
```

### Eliminar una tarea
```bash
curl -X DELETE http://localhost:8080/api/todos/1
```
