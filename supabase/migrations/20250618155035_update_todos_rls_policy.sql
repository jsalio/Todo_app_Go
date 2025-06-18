-- Eliminar la política existente
DROP POLICY IF EXISTS "Permitir acceso autenticado" ON todos;

-- Crear una nueva política que permita acceso público
CREATE POLICY "Permitir acceso público" ON todos
    FOR ALL TO public USING (true);