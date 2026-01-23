package handler

import (
	"net/http"
	"strings"
	"ulasin-backend/model"
	"ulasin-backend/repository"
	// utils
	"ulasin-backend/utils"

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

// Section Login
// inisialisasi struct request login
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Login Header
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
			"error": err.Error(),
		})
	}

	// trim
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email dan password wajib diisi.",
		})
	}

	// cari user
	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Email atau password salah",
			"error": err.Error(),
		})
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Email atau password salah",
			"error": err.Error(),
		})
	}

	// generate json web token (jwt)
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal generate token",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token": token,
		"user": fiber.Map{
			"id": user.ID,
			"nama": user.Nama,
			"email": user.Email,
			"role": user.Role,
		},
	})
}

// End Section Login