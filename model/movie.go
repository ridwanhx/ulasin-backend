package model

import "time"

type Movie struct {
	ID uint `gorm:"primaryKey"`
	Judul string `gorm:"type:varchar(150);not null"`
	Poster string `gorm:"type:text;"`
	Sutradara string `gorm:"type:varchar(100)"`
	Genre string `gorm:"varchar(50)"`
	TahunRilis int
	Sinopsis string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// inisialisasi foreign key (relasi dengan Reviews)
	Reviews []Review `gorm:"foreignKey:MovieID"`
}