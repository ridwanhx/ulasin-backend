// Setup Fiber Server
package main

import (
	"log"
	"os"
	"ulasin-backend/config"
	"ulasin-backend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// konek ke DB
	if err := config.InitDatabase(); err != nil {
		log.Fatal("Gagal terhubung ke database", err)
	}
	log.Println("Berhasil terhubung ke Database")

	// jalankan Auto Migration
	if err := config.Migrate(); err != nil {
		log.Fatal("Gagal migrate database:", err)
	}
	log.Println("Migrasi database berhasil")

	// instansiasi object fiber baru
	app := fiber.New()

	// jalankan middleware
	app.Use(cors.New(config.SetupCORS()))

	// panggil router
	router.SetupRoutes(app)
	
	// route untuk testing koneksi ke db
	app.Get("/health", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app": os.Getenv("APP_NAME"),
		})
	})

	// inisialisasi port dari .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}