# Go Todo App with Supabase

Una aplicación de línea de comandos (CLI) para gestionar tareas, construida en Go y utilizando Supabase como backend.

## Características

- ✅ Crear nuevas tareas
- 📋 Listar todas las tareas
- ✏️ Actualizar tareas existentes
- ❌ Eliminar tareas
- 💾 Almacenamiento en la nube con Supabase

## Requisitos Previos

- [Go](https://golang.org/dl/) 1.16 o superior
- Cuenta en [Supabase](https://supabase.com/)
- Base de datos configurada en Supabase con una tabla `todos`

## Configuración

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/sample-todo-app.git
   cd sample-todo-app
   ```

2. Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:
   ```
   SUPABASE_URL=tu_url_de_supabase
   SUPABASE_KEY=tu_clave_de_supabase
   ```

3. Instala las dependencias:
   ```bash
   go mod download
   ```
4. Instala las dependencias de node:
   ```bash
   npm i
   ```

## Estructura de la Base de Datos

Crea una tabla `todos` en Supabase con la siguiente estructura:

```sql
create table todos (
  id bigint generated by default as identity primary key,
  task text not null,
  is_complete boolean default false,
  create_at timestamp with time zone default timezone('utc'::text, now()) not null
);
```

### link project to supabase

Initialize session on supabase :

   ```bash
   npx supabase login
   ```

For link project to Supabase run :
   ```bash
   npx supabase link --project-ref=you-project-id
   ```

Run the comamd for create schema on supabase:
```bash
   npx supabase db push
```

## Create migrations

para crear una migración de tabla run:
   ```bash
   npx supabase migration new create_new_table
   ```
esto creara un archivo sql en `supabase/migrations/YYYYMMDDHHMMSS_create_new_table.sql`. Editelo para que coincida con la structura.ejemplo:
   ```sql
   -- Crear tabla todos
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    is_complete BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Habilitar Row-Level Security (RLS)
ALTER TABLE todos ENABLE ROW LEVEL SECURITY;

-- Crear política para acceso autenticado
CREATE POLICY "Permitir acceso autenticado" ON todos
    FOR ALL TO authenticated USING (true);
   ```

## Uso

### Listar tareas
```bash
go run main.go list
```

### Agregar una nueva tarea
```bash
go run main.go add "Mi nueva tarea"
```

### Marcar una tarea como completada
```bash
go run main.go update <id> --complete
```

### Actualizar el texto de una tarea
```bash
go run main.go update <id> --task "Nuevo texto de la tarea"
```

### Eliminar una tarea
```bash
go run main.go delete <id>
```

## Construir el binario

Para crear un ejecutable del CLI:

```bash
go build -o todo-cli
```
Or compile and run 
```bash
go build -o todo-app && ./todo-app serve
```

Luego puedes usarlo así:

```bash
./todo-cli list
```

## Contribuir

Las contribuciones son bienvenidas. Por favor, abre un issue primero para discutir los cambios propuestos.

## Licencia

Este proyecto está bajo la [Licencia MIT](LICENSE).
