# ProgramacionGrupoA1

## Sistema de Gestión de Libros Electrónicos (Biblioteca Virtual)

Proyecto desarrollado en Golang para la gestión de libros electrónicos, permitiendo administrar usuarios, consultar libros, realizar búsquedas y gestionar lecturas dentro de una biblioteca virtual.

Desarrollado como parte de la asignatura Programación Orientada a Objetos de la Universidad Internacional del Ecuador.
## Repositorio

Nombre del repositorio: ProgramacionGrupoA1

## Docente

Ing. Héctor Guillermo Ávalos Silva

## Universidad

Universidad Internacional del Ecuador
Carrera de Ingeniería en Ciberseguridad

# 📚 Biblioteca Virtual

## Descripción del Proyecto

Biblioteca Virtual es una aplicación desarrollada en Golang que permite gestionar libros electrónicos de manera sencilla y organizada. El sistema facilita la administración de usuarios, la consulta de libros, la búsqueda por diferentes criterios y el seguimiento de lecturas realizadas por los usuarios.

Este proyecto ha sido desarrollado como parte de la asignatura **Programación Orientada a Objetos** de la Universidad Internacional del Ecuador.

---

## Objetivo General

Desarrollar una aplicación orientada a la gestión de una biblioteca virtual que permita administrar libros electrónicos, facilitar la búsqueda de información y registrar el historial de lectura de los usuarios mediante una estructura simple y organizada.

---

## Objetivos Específicos

- Permitir el registro e inicio de sesión de usuarios.
- Gestionar un catálogo de libros electrónicos.
- Facilitar la búsqueda de libros por diferentes criterios.
- Registrar el progreso de lectura de cada usuario.
- Administrar la información almacenada en el sistema.
- Permitir la descarga de libros digitales disponibles.

---

## Módulos del Sistema

### 🔐 Módulo 1: Autenticación

Funciones:

- Registro de usuarios.
- Inicio de sesión.
- Validación de credenciales.
- Cierre de sesión.
- Acceso para administradores.

### 📖 Módulo 2: Catálogo de Libros

Funciones:

- Visualizar libros disponibles.
- Buscar libros por título.
- Buscar libros por autor.
- Buscar libros por género.
- Visualizar sinopsis.
- Descargar libros.

### 👤 Módulo 3: Cuenta de Usuario

Funciones:

- Ver perfil del usuario.
- Consultar historial de lectura.
- Ver libros descargados.
- Eliminar cuenta.

### ⚙️ Módulo 4: Administración

Funciones:

- Agregar libros.
- Editar libros.
- Eliminar libros.
- Consultar usuarios registrados.
- Gestionar soporte.

---

## Tecnologías Utilizadas

- Go (Golang)
- JSON
- GitHub
- Visual Studio Code
- Gorilla Mux

---

## Arquitectura del Proyecto

```
biblioteca-virtual/

├── main.go
│
├── models/
│   ├── usuario.go
│   ├── libro.go
│   └── perfil.go
│
├── handlers/
│   ├── auth_handler.go
│   ├── usuario_handler.go
│   └── libro_handler.go
│
├── services/
│   ├── auth_service.go
│   ├── usuario_service.go
│   └── libro_service.go
│
├── routes/
│   └── routes.go
│
├── data/
│   ├── usuarios.json
│   └── libros.json
│
└── go.mod
```

---

## Temas de Golang Aplicados

| Tema | Aplicación |
|--------|------------|
| Structs | Modelado de entidades |
| Slices | Gestión de listas |
| Maps | Búsquedas rápidas |
| Functions | Procesamiento de datos |
| Closures | Validaciones |
| Interfaces | Comunicación entre módulos |
| Packages | Organización del sistema |
| JSON | Persistencia de datos |
| HTTP | Comunicación cliente-servidor |

---
