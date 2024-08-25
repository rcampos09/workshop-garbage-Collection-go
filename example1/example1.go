package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var data [][]byte

// allocateMemory maneja las solicitudes para asignar memoria
func allocateMemory(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 10000; i++ {
		data = append(data, make([]byte, 1024*1024)) // Asigna 1MB de memoria y reténlo
	}
	fmt.Fprintf(w, "Allocated 10GB of memory!\n")
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
