package middleware

import (
	"strings"
	"ulasin-backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Authorization header tidak ditemukan",
			})
		}

		// format Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Format Authorization harus Bearer <token>",
			})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func (token *jwt.Token) (interface{}, error) {
			// pastikan signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token tidak valid atau sudah kadaluarsa",
				"err": err.Error(),
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Gagal membaca token",
			})
		}

		// ambil data dari token
		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)

		// simpan ke context
		c.Locals("user_id", userID)
		c.Locals("role", role)

		return c.Next()
	}
}