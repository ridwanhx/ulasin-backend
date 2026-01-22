// Setup Fiber Server
package main

import (
	"log"
	"os"
	"ulasin-backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := config.InitDatabase(); err != nil {
		log.Fatal("Gagal terhubung ke database", err)
	}
	log.Println("Berhasil terhubung ke Database")

	app := fiber.New()

	app.Use(cors.New(config.SetupCORS()))

	app.Get("/health", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app": os.Getenv("APP_NAME"),
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}