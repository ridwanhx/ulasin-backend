package repository

import (
	"ulasin-backend/config"
	"ulasin-backend/model"

	"gorm.io/gorm"
)

type MovieRepository struct{}

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
	err := config.DB.Find(&movies).Error
	return movies, err
}

// Get By ID
func (r *MovieRepository) FindByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	err := config.DB.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// Update
func (r *MovieRepository) Update(movie *model.Movie) error {
	return config.DB.Save(movie).Error
}

// Delete
func (r *MovieRepository) Delete(id uint) error {
	result := config.DB.Delete(&model.Movie{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}