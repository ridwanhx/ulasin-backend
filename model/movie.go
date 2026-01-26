package model

import "time"

type Movie struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Judul string `gorm:"type:varchar(150);not null" json:"judul"`
	Poster string `gorm:"type:text;" json:"poster"`
	Sutradara string `gorm:"type:varchar(100)" json:"sutradara"`
	Genre string `gorm:"varchar(50)" json:"genre"`
	TahunRilis int `json:"tahun_rilis"`
	Sinopsis string `gorm:"type:text" json:"sinopsis"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// computed fields (Fields yang tidak akan di migrasi / pseudo column)
	AverageRating float64 `gorm:"-" json:"average_rating"`
	TotalReviews int64 `gorm:"-" json:"total_reviews"`

	// inisialisasi foreign key (relasi dengan Reviews)
	Reviews []Review `gorm:"foreignKey:MovieID"`
}