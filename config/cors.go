package config

import "github.com/gofiber/fiber/v2/middleware/cors"

// setup cors
func SetupCORS() cors.Config {
	return cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}
}