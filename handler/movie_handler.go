package handler

import (
	"errors"
	"strconv"
	"ulasin-backend/model"
	"ulasin-backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MovieHandler struct {
	movieRepo *repository.MovieRepository
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler {
		movieRepo: repository.NewMovieRepository(),
	}
}

// create movie (admin only)
func (h *MovieHandler) Create(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak",
		})
	}

	var movie model.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data tidak valid",
			"error": err.Error(),
		})
	}

	if movie.Judul == "" ||
		movie.Sutradara == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Judul dan Sutradara wajib diisi",
		})
	}

	if err := h.movieRepo.Create(&movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan data movie",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Movie berhasil ditambahkan",
		"data": movie,
	})
}

// Get All Movies (Public)
func (h *MovieHandler) FindAll(c *fiber.Ctx) error {
	movies, err := h.movieRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data movie",
		})
	}

	return c.JSON(fiber.Map{
		"data": movies,
	})
}

// Get Movie by ID (Public)
func (h *MovieHandler) FindByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak valid",
			"error": err.Error(),
		})
	}

	movie, err := h.movieRepo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Movie tidak ditemukan",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": movie,
	})
}

// Update Movie (admin)
func (h *MovieHandler) Update(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak",
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	movie, err := h.movieRepo.FindByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data movie tidak ditemukan",
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
			"error": err.Error(),
		})
	}

	if err := h.movieRepo.Update(movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal update movie",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Movie berhasil diupdate",
		"data": movie,
	})
}

// Delete movie (admin)
func (h *MovieHandler) Delete(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID Tidak valid",
			"error": err.Error(),
		})
	}

	err = h.movieRepo.Delete(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Movie tidak ditemukan",
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus movie",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Movie berhasil dihapus",
	})
}