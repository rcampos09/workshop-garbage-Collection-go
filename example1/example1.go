package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Importa el paquete pprof para habilitar el perfilado del rendimiento
)

// Declaración de una variable global 'data' que almacenará bloques de memoria
var data [][]byte

// allocateMemory maneja las solicitudes HTTP para asignar memoria
func allocateMemory(w http.ResponseWriter, r *http.Request) {
	// Asigna 1MB de memoria 10,000 veces y retén cada bloque en la variable 'data'
	for i := 0; i < 10000; i++ {
		data = append(data, make([]byte, 1024*1024)) // Asigna y retén 1MB por iteración
	}
	// Envía una respuesta indicando que se han asignado 10GB de memoria
	fmt.Fprintf(w, "Allocated 10GB of memory!\n")
}

func main() {
	// Inicia un servidor adicional para pprof en el puerto 6060
	go func() {
		log.Println("Starting pprof server on :6060") // Muestra un mensaje de inicio del servidor pprof
		err := http.ListenAndServe(":6060", nil)      // Escucha y sirve el servidor pprof en el puerto 6060
		if err != nil {                               // Si ocurre un error, lo registra y termina el programa
			log.Fatalf("pprof server failed: %v", err)
		}
	}()

	// Configura y empieza el servidor principal en el puerto 8080
	http.HandleFunc("/allocate", allocateMemory) // Asocia la ruta '/allocate' con la función allocateMemory
	fmt.Println("Starting main server on :8080") // Imprime un mensaje indicando que el servidor principal está iniciando
	err := http.ListenAndServe(":8080", nil)     // Escucha y sirve el servidor principal en el puerto 8080
	if err != nil {                              // Si ocurre un error, lo registra y termina el programa
		log.Fatalf("main server failed: %v", err)
	}
}
