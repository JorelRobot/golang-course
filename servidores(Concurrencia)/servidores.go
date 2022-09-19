package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de transcurrido %s\n", tiempoTranscurrido)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Printf("El servidor %s se encuentra caido\n", servidor)
	} else {
		fmt.Printf("El servidor %s esta funcionando normalmente\n", servidor)
	}
}
