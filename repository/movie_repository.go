package repository

import (
	"ulasin-backend/config"
	"ulasin-backend/model"

	"gorm.io/gorm"
)

type MovieRepository struct{}

type RatingResult struct {
	AverageRating float64
	TotalReviews int64
}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

// Create
func (r *MovieRepository) Create(movie *model.Movie) error {
	return config.DB.Create(movie).Error
}

// Get All
func (r *MovieRepository) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	err := config.DB.
		Preload("Reviews").
		Preload("Reviews.User").
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	for i := range movies {
		_ = r.fillRating(&movies[i])
	}
	return movies, nil
}

// Get By ID
func (r *MovieRepository) FindByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	err := config.DB.
		Preload("Reviews").
		Preload("Reviews.User").
		First(&movie, id).Error
		
	if err != nil {
		return nil, err
	}

	// add fillrating
	_ = r.fillRating(&movie)
	return &movie, nil
}

// Update
func (r *MovieRepository) Update(movie *model.Movie) error {
	return config.DB.Save(movie).Error
}

// Delete
func (r *MovieRepository) Delete(id uint) error {
	// Gunakan Transaction untuk memastikan kedua proses berhasil
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Hapus review yang berkaitan dulu
		if err := tx.Where("movie_id = ?", id).Delete(&model.Review{}).Error; err != nil {
			return err
		}

		// 2. Baru hapus Movie-nya
		result := tx.Delete(&model.Movie{}, id)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})
}

// Helper Function
func (r *MovieRepository) fillRating(movie *model.Movie) error {
	var result RatingResult

	err := config.DB.Model(&model.Review{}).Select("AVG(skor) as average_rating, COUNT(*) as total_reviews").Where("movie_id = ?", movie.ID).Scan(&result).Error

	if err != nil {
		return err
	}

	movie.AverageRating = result.AverageRating
	movie.TotalReviews = result.TotalReviews

	return nil
}