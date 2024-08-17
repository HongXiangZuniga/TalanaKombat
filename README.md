# TalanaKombat

Esta es una prueba técnica desarrollada para Talana, utilizando Go y la arquitectura hexagonal. El proyecto implementa un juego de combate simple, donde los jugadores pueden realizar movimientos y combos predefinidos. La arquitectura hexagonal permite dividir el proyecto en tres capas principales: 

-**Game**: Encargada del bucle principal del juego (game loop).
-**Relator**: Responsable de los mensajes de resultado.
-**Fight**: Gestiona la validación de movimientos y combos.

Este documento explica cómo configurar y ejecutar el proyecto, tanto en entorno local como dentro de un contenedor Docker.

### Configuración

Antes de ejecutar la aplicación, asegúrate de crear un archivo `.env` con las siguientes variables de entorno:

```bash
ENV= [dev|prod]  # Define el entorno de ejecución (desarrollo o producción).
PORT= [Port for your app]  # Especifica el puerto en el que correrá la aplicación.
```
#### Ejemplo
```bash
ENV=dev
PORT=8080
```

Tienes una collecion de postman con los request asociados a los enunciados de la prueba. Se corrigio el primero que presentaba una diferencia en los mensajes de salida con el input de ejemplo.

```bash
collection.json
```


### Instalación de Dependencias

Para instalar las dependencias necesarias para el proyecto, ejecuta el siguiente comando:

```bash
make install  # Instala todas las dependencias especificadas en el go.mod
```

### Ejecutar proyecto en local 
 
```bash
Make run # Corre el proyecto en local 
```

### Para crear Docker

```bash
Make docker-build # Crea un docker del proyecto 
```

### Para correr Docker

```bash
Make docker-run #Ejecuta el docker creado
```

### Organización del Proyecto

La estructura del proyecto sigue la arquitectura hexagonal, separando las responsabilidades en diferentes capas:

```bash
.
├── Makefile        # Definición de comandos para la gestión del proyecto
├── README.md       # Este archivo
├── build
│   └── Dockerfile  # Archivo para la construcción de la imagen de Docker
├── cmd
│   └── api
│       └── main.go  # Punto de entrada de la aplicación
├── collection.json  # Colección de Postman para pruebas
├── go.mod          # Módulo de Go con dependencias
├── go.sum          # Suma de verificación de las dependencias
└── internal        # Lógica interna de la aplicación
    ├── adapters    # Adaptadores para la interfaz HTTP y repositorios
    │   ├── http    # Manejo de rutas y controladores
    │   │   └── handler
    │   │       └── game.go  # Controlador para el game loop
    │   └── repository
    │       └── fight.go  # Implementación del repositorio para combates
    └── core        # Núcleo de la lógica de negocio
        └── domain  # Definición de dominios y servicios
            ├── fight
            │   ├── model.go       # Modelo de datos para combates
            │   ├── repository.go  # Interfaz del repositorio de combates
            │   └── service.go     # Lógica de negocio para combates
            ├── game
            │   ├── model.go       # Modelo de datos para el juego
            │   └── service.go     # Lógica de negocio para el juego
            └── relator
                └── service.go     # Servicio para generación de mensajes
```

## Conclusión

¡Gracias por tomarte el tiempo para revisar esta prueba técnica! Espero que encuentres el código y la arquitectura claros y bien estructurados. Si tienes alguna pregunta o necesitas más información, no dudes en contactarme.