# Control Escolar

Este proyecto es un sistema web de control escolar desarrollado en Go (Golang) usando el framework Gin y GORM para la gestión de base de datos SQLite. Permite gestionar estudiantes, materias, inscripciones y calificaciones, asegurando integridad referencial y relaciones correctas entre las tablas.

## Características principales
- Alta, baja y modificación de estudiantes y materias.
- Inscripción de estudiantes en materias (relación muchos a muchos).
- Asignación de calificaciones solo a estudiantes inscritos en materias.
- Consulta y eliminación de calificaciones.
- Interfaz web sencilla para todas las operaciones.

---

# Tutorial de instalación y uso

## 1. Requisitos previos
- Tener instalado [Go](https://go.dev/dl/) (versión 1.18 o superior recomendada).
- Tener instalado [Git](https://git-scm.com/downloads).
- Sistema operativo Windows, Linux o MacOS.

## 2. Clonar el repositorio

Abre una terminal (PowerShell en Windows, git bash o git cmd) y ejecuta:

```powershell
git clone <URL_DEL_REPOSITORIO>
cd <CARPETA_DEL_REPOSITORIO>/control_escolar
```

## 3. Instalar dependencias

Dentro de la carpeta `control_escolar`, ejecuta:

```powershell
go mod tidy
```

Esto descargará todas las librerías necesarias (Gin, GORM, SQLite, etc).

## 4. Ejecutar el programa

Asegúrate de estar en la carpeta `control_escolar` y ejecuta:

```powershell
go run main.go
```

Verás un mensaje como:

```
Servidor escuchando en el puerto 3000
```

## 5. Acceder a la aplicación web

Abre tu navegador y entra a:

```
http://localhost:3000/
```

## 6. Uso de la aplicación

- **Agregar estudiantes y materias:** Usa los formularios en la web.
- **Inscribir estudiantes en materias:** Usa el formulario "Inscribir estudiante en materia".
- **Agregar calificaciones:** Solo podrás calificar materias en las que el estudiante esté inscrito.
- **Buscar, editar y eliminar:** Usa los formularios y botones correspondientes en la interfaz.

## 7. Estructura de la base de datos

- `students`: Almacena los estudiantes.
- `subjects`: Almacena las materias.
- `student_subjects`: Relación muchos a muchos entre estudiantes y materias.
- `grades`: Calificaciones, asociadas a un estudiante y una materia.

## 8. Notas importantes

- Si cambias los modelos, elimina el archivo `control_escolar.db` para regenerar la base de datos.
  esto porque Las migraciones automáticas de GORM (AutoMigrate) no siempre manejan correctamente todos los cambios en la estructura, especialmente en SQLite
Pueden quedar datos inconsistentes o con estructura antigua que no coincida con los nuevos modelos
SQLite no soporta todas las operaciones ALTER TABLE que sí soportan otras bases de datos y por esto
Pueden persistir problemas de bloqueo de la base de datos debido a transacciones incompletas
- El sistema valida que solo se puedan asignar calificaciones a estudiantes inscritos en la materia.
- El puerto por defecto es el 3000, puedes cambiarlo en `main.go`.

## 9. Problemas comunes

- **Error de dependencias:** Ejecuta `go mod tidy`.
- **Puerto ocupado:** Cambia el puerto en `main.go`.
- **No se ven materias en el selector de calificaciones:** Inscribe primero al estudiante en la materia.

---

# Créditos
Desarrollado por "Meg" Raimond, 2025.
