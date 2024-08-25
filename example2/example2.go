package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

var mu sync.Mutex

// allocateMemory maneja las solicitudes para asignar y liberar memoria
func allocateMemory(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1024*1024) // Asignar 1MB de memoria, pero sin retener
	}

	// Forzar la recolección de basura para liberar la memoria inmediatamente
	runtime.GC()

	fmt.Fprintf(w, "Allocated and freed 10GB of memory!\n")
}

func main() {
	// Inicia el servidor pprof en el puerto 6060
	go func() {
		log.Println("Starting pprof server on :6060")
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			log.Fatalf("pprof server failed: %v", err)
		}
	}()

	// Configura el servidor principal en el puerto 8080
	http.HandleFunc("/allocate", allocateMemory)
	fmt.Println("Starting main server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("main server failed: %v", err)
	}
}
