package main

import (
	"log"

	"github.com/Fran313/retailBrain/api"
	"github.com/Fran313/retailBrain/config"
	"github.com/Fran313/retailBrain/internal/vectorstore"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar conexión a Postgres
	if err := config.InitDB(); err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	// Inicializar Qdrant
	vectorstore.InitQdrant()

	// Iniciar servidor
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080")
}
