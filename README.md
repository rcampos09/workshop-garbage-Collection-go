# Workshop Garbage Collection Go (Example)

Este proyecto contiene dos ejemplos de cómo se maneja el Garbage Collection (GC) en Go. Cada ejemplo está configurado para ejecutarse en un contenedor Docker y se puede usar para analizar cómo el Garbage Collection afecta el rendimiento de las aplicaciones Go.


- **example1/**: Contiene un ejemplo donde la memoria asignada no es liberada, lo que puede llevar a un uso excesivo de memoria.
- **example2/**: Contiene un ejemplo optimizado donde se fuerza la liberación de memoria usando `runtime.GC()`.

## Requisitos Previos

- Docker
- Go (si deseas ejecutar el código fuera de Docker)

## Cómo Construir y Ejecutar los Contenedores

### Example 1

1. Navega al directorio `example1`:

    ```bash
    cd example1
    ```

2. Construye el contenedor Docker:

    ```bash
    docker build -t gc-example1 .
    ```

3. Ejecuta el contenedor:

    ```bash
    docker run -d -p 8080:8080 -p 6160:6060 --name gc-example1-container gc-example1
    ```

4. Accede al servicio desde el navegador o usando `curl`:

    ```bash
    curl http://localhost:8080/allocate
    ```

5. Analiza el GC y otros perfiles:

    ```bash
    go tool pprof http://localhost:6160/debug/pprof/profile
    ```

### Example 2

1. Navega al directorio `example2`:

    ```bash
    cd example2
    ```

2. Construye el contenedor Docker:

    ```bash
    docker build -t gc-example2 .
    ```

3. Ejecuta el contenedor:

    ```bash
    docker run -d -p 8081:8080 -p 6260:6060 --name gc-example2-container gc-example2
    ```

4. Accede al servicio desde el navegador o usando `curl`:

    ```bash
    curl http://localhost:8081/allocate
    ```

5. Analiza el GC y otros perfiles:

    ```bash
    go tool pprof http://localhost:6260/debug/pprof/profile
    ```

## Análisis de los Perfiles

Puedes utilizar `go tool pprof` para analizar el uso de CPU y memoria de cada contenedor:

- **Perfil de CPU**:

    ```bash
    go tool pprof http://localhost:6160/debug/pprof/profile?seconds=30
    ```

- **Perfil de Memoria**:

    ```bash
    go tool pprof http://localhost:6160/debug/pprof/heap
    ```

- **Trazado en vivo**:

    ```bash
    go tool pprof http://localhost:6160/debug/pprof/trace
    ```

Repite los mismos comandos para `example2`, cambiando los puertos según corresponda (`6260` en lugar de `6160`).

## Conclusión

Este proyecto proporciona un entorno simple para ver cómo el Garbage Collection afecta las aplicaciones Go, permitiendo comparaciones directas entre un código que retiene memoria y uno que libera activamente memoria. Puedes extender este proyecto para explorar más optimizaciones y análisis de rendimiento en Go.


### Propiedad y Derechos de Autor
Este código es propiedad de Rodrigo Campos (@Dontester). Todos los derechos de autor están reservados por Rodrigo Campos.

© Rodrigo Campos (@Dontester)