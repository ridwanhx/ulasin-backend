package handler

import (
	"net/http"
	"ulasin-backend/model"
	"ulasin-backend/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userRepo: repository.NewUserRepository(),
	}
}

type RegisterRequest struct {
	Nama string `json:"nama"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	// tangkap request
	var req RegisterRequest
	// parsing nilai hasil request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
			"error": err.Error(),
		})
	}

	// pastikan request field tidak kosong
	if req.Nama == "" ||
		req.Email == "" ||
		req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Field nama, email, dan password wajib diisi.",
		})
	}

	// pastikan panjang pasword >= 8 karakter
	if len(req.Password) < 8 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password minimal 8 karakter",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error": err.Error(),
		})
	}

	user := model.User {
		Nama: req.Nama,
		Email: req.Email,
		Password: string(hashedPassword),
		Role: "user",
	}

	// cek apakah email sudah terdaftar
	if err := h.userRepo.Create(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email sudah terdaftar",
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"data": fiber.Map{
			"id": user.ID,
			"nama": user.Nama,
			"email": user.Email,
			"role": user.Role,
		},
	})
}