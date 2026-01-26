package router

import (
	"ulasin-backend/handler"
	"ulasin-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// auth handler
	authHandler := handler.NewAuthHandler()

	// movie handler
	movieHandler := handler.NewMovieHandler()

	// review handler
	reviewHandler := handler.NewReviewHandler()

	api := app.Group("/api")

	// -------------
	// Public Routes
	// -------------

	// authentication routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// movie routes
	api.Get("/movies", movieHandler.FindAll)
	api.Get("/movies/:id", movieHandler.FindByID)

	// review routes
	api.Get("/movies/:movieId/reviews", reviewHandler.FindByMovie)

	// ----------------
	// Protected Routes
	// ----------------
	protected := api.Group("", middleware.JWTProtected())

	// movie routes
	protected.Post("/movies", movieHandler.Create)
	protected.Put("/movies/:id", movieHandler.Update)
	protected.Delete("/movies/:id", movieHandler.Delete)

	// review routes
	protected.Post("/movies/:movieId/reviews", reviewHandler.Create)
	protected.Put("/reviews/:id", reviewHandler.Update)
	protected.Delete("/reviews/:id", reviewHandler.Delete)
}