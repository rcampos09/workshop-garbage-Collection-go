package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Importa el paquete pprof para habilitar el perfilado del rendimiento
	"runtime"          // Importa el paquete runtime para interactuar con el recolector de basura
	"sync"             // Importa el paquete sync para utilizar un mutex (exclusión mutua)
)

// Declaración de un mutex para sincronizar el acceso a la función allocateMemory
var mu sync.Mutex

// allocateMemory maneja las solicitudes HTTP para asignar y liberar memoria
func allocateMemory(w http.ResponseWriter, r *http.Request) {
	// Bloquea el mutex al entrar y lo desbloquea al salir de la función
	mu.Lock()
	defer mu.Unlock()

	// Asigna 1MB de memoria 10,000 veces pero sin retenerlo
	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1024*1024) // Asigna 1MB por iteración sin guardar referencia
	}

	// Forza la recolección de basura para liberar la memoria inmediatamente
	runtime.GC()

	// Envía una respuesta indicando que se han asignado y liberado 10GB de memoria
	fmt.Fprintf(w, "Allocated and freed 10GB of memory!\n")
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
