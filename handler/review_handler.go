package handler

import (
	"strconv"
	"ulasin-backend/model"
	"ulasin-backend/repository"

	"github.com/gofiber/fiber/v2"
)

// inisialisasi struct
type ReviewHandler struct {
	reviewRepo *repository.ReviewRepository
}

// inisialisasi method handler
func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{
		reviewRepo: repository.NewReviewRepository(),
	}
}

// inisialisasi method create review
func (h *ReviewHandler) Create(c *fiber.Ctx) error {
	// ambil nilai user id (foreign key)
	userID := c.Locals("user_id").(uint)

	// deklarasi var untuk menyimpan movie id, dan err untuk menyimpan nilai kembalian jika nanti terjadi error
	movieID, err := strconv.Atoi(c.Params("movieId"))
	// jika nilai id tidak valid
	if err != nil {
		// kembalikan status bad request
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Movie ID tidak valid",
			"error": err.Error(),
		})
	}

	// inisialisasi variabel untuk menyimpan request
	var req struct {
		Skor int `json:"skor"`
		Komentar string `json:"komentar"`
	}

	// parsing nilai request yang masuk
	if err := c.BodyParser(&req); err != nil {
		// kembalikan status bad request / 400 jika error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// beri pesan kesalahan
			"message": "Data tidak valid",
			"error": err.Error(),
		})
	}

	// atur supaya request untuk skor yang dikirimkan tidak boleh kurang/lebih dari 5 (1 - 5)
	if req.Skor < 1 || req.Skor > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Skor harus antara 1 sampai 5",
		})
	}

	if _, err := h.reviewRepo.FindByUserAndMovie(userID, uint(movieID));err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamu sudah memberi review untuk movie ini",
		})
	}

	// simpan setiap nilai kedalam struct
	review := model.Review{
		UserID: userID,
		MovieID: uint(movieID),
		Skor: req.Skor,
		Komentar: req.Komentar,
	}

	// jika terjadi error pada server (diluar kesalahan user)
	if err := h.reviewRepo.Create(&review); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan review",
			"error": err.Error(),
		})
	}

	createdReview, _ := h.reviewRepo.FindByIDWithRelations(review.ID)

	// kembalikan status review berhasil dibuat
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Review berhasil ditambahkan",
		"data": createdReview,
	})
}

// Get Reviews by Movie (Public)
func (h *ReviewHandler) FindByMovie(c *fiber.Ctx) error {
	movieID, err := strconv.Atoi(c.Params("movieId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Movie ID tidak valid",
			"error": err.Error(),
		})
	}

	reviews, err := h.reviewRepo.FindByMovieID(uint(movieID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil review",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": reviews,
	})
}

// Update Review (owner only)
func (h *ReviewHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, _ := strconv.Atoi(c.Params("id"))
	review, err := h.reviewRepo.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Review tidak ditemukan",
		})
	}

	if review.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak",
		})
	}

	if err := c.BodyParser(review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data tidak valid",
			"error": err.Error(),
		})
	}

	if err := h.reviewRepo.Update(review); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal update review",
		})
	}

	updatedReview, _ := h.reviewRepo.FindByIDWithRelations(review.ID)

	return c.JSON(fiber.Map{
		"message": "Review berhasil diupdate",
		"data": updatedReview,
	})
}

// Delete Review (owner atau admin)
func (h *ReviewHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	id, _ := strconv.Atoi(c.Params("id"))
	review, err := h.reviewRepo.FindByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Review tidak ditemukan",
			"error": err.Error(),
		})
	}

	if review.UserID != userID && role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Akses ditolak",
		})
	}

	if err := h.reviewRepo.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus review",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Review berhasil dihapus",
	})
}