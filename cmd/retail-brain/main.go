package main

import (
	"log"
	"os"

	"github.com/Fran313/retailBrain/api"
	"github.com/Fran313/retailBrain/internal/database"
	"github.com/Fran313/retailBrain/internal/vectorstore"
	"github.com/gin-gonic/gin"

	// Importa el paquete de migrate
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://admin:admin123@localhost:5432/retailBrain?sslmode=disable"
	}
	m, err := migrate.New(
		"file://./db/migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("✅ Database migrations applied")
}

func main() {
	// Inicializar conexión a Postgres
	if err := database.InitDB(); err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	runMigrations()

	// Inicializar Qdrant
	vectorstore.InitQdrant()

	// Iniciar servidor
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080")
}
