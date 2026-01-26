package repository

import (
	"ulasin-backend/config"
	"ulasin-backend/model"

	"gorm.io/gorm"
)

type ReviewRepository struct{}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{}
}

func (r *ReviewRepository) Create(review *model.Review) error {
	return config.DB.Create(review).Error
}

func (r *ReviewRepository) FindByMovieID(movieID uint) ([]model.Review, error) {
	var reviews []model.Review
	err := config.DB.
		Preload("User").
		Where("movie_id = ?", movieID).
		Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) FindByID(id uint) (*model.Review, error) {
	var review model.Review
	err := config.DB.
		Preload("User").
		First(&review, id).Error

	return &review, err
}

func (r *ReviewRepository) Update(review *model.Review) error {
	return config.DB.Save(review).Error
}

func (r *ReviewRepository) Delete(id uint) error {
	result := config.DB.Delete(&model.Review{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Validasi: 1 user hanya boleh 1 review per movie
func (r *ReviewRepository) FindByUserAndMovie(userID, movieID uint) (*model.Review, error) {
	var review model.Review
	err := config.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).First(&review).Error

	if err != nil {
		return nil, err
	}
	return &review, nil
}