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