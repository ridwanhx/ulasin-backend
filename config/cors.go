package config

import "github.com/gofiber/fiber/v2/middleware/cors"

// setup cors
func SetupCORS() cors.Config {
	return cors.Config{
		AllowOrigins: "http://localhost:5173,https://ulasin-frontend.vercel.app,https://ulasin-frontend-bn4z8eyc7-muhamad-ridwans-projects-fddf9a10.vercel.app",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}
}