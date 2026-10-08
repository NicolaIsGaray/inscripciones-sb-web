package main

import (
	"log"
	"net/http"
	"os"

	"servidor-angular/config"
	"servidor-angular/router"
)

func main() {
	// Conectar a MongoDB
	if err := config.Connect(); err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer config.Disconnect()

	// Configurar router
	r := router.NewRouter()

	// Aplicar middleware de CORS
	handler := router.CorsMiddleware(r)

	// Obtener puerto desde variable de entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	log.Printf("Servidor API iniciado en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
