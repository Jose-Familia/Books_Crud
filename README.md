﻿# Books_Crud [API CRUD de Libros en Golang]

## Introducción
Este proyecto implementa una API RESTful en Golang para gestionar una colección de libros, permitiendo las operaciones CRUD (Crear, Leer, Actualizar y Eliminar).

## Diagrama de la Base de Datos
[Insertar aquí la imagen de tu esquema de base de datos]

* **Descripción breve del esquema:**
  - Explica qué tablas componen la base de datos y cómo se relacionan entre sí.
  - Menciona los campos clave de cada tabla (ID, título, autor, etc.).
  - Indica el tipo de datos utilizado para cada campo.

## Funcionamiento
### Pre-requisitos
* **Golang:** Asegúrate de tener instalado Golang en tu sistema. Descarga la última versión desde https://golang.org/dl/
* **Gestor de base de datos:** Instala y configura el gestor de base de datos que estás utilizando (MySQL, PostgreSQL, etc.).

## Endpoints
- GET /books: Obtiene una lista de todos los libros.
- POST /books: Crea un nuevo libro.

### Cuerpo de la solicitud: JSON con los datos del libro (título, autor, etc.).
- GET /books/:id: Obtiene un libro específico por su ID.
- PUT /books/:id: Actualiza un libro existente.

### Cuerpo de la solicitud: JSON con los nuevos datos del libro.
- DELETE /books/:id: Elimina un libro.

## Estructura del Código
- database: Contiene el código para interactuar con la base de datos.
- handlers: Contiene los manejadores de las solicitudes HTTP.
- models: Define las estructuras de datos que representan los libros.

## Base de datos

![image](https://github.com/user-attachments/assets/bd6e4f20-336d-48e7-996c-3a6666b0cd6f)


