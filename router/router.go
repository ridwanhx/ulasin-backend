package router

import (
	"ulasin-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	authHandler := handler.NewAuthHandler()

	api := app.Group("/api")

	api.Post("/register", authHandler.Register)
}