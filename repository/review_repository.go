package repository

import (
	"ulasin-backend/config"
	"ulasin-backend/model"
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
	err := config.DB.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(review *model.Review) error {
	return config.DB.Save(review).Error
}

func (r *ReviewRepository) Delete(id uint) error {
	return config.DB.Delete(&model.Review{}, id).Error
}