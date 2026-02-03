package model

import "time"

type Review struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"not null"`
	MovieID uint `gorm:"not null"`
	Skor int `gorm:"not null"`
	Komentar string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// inisialisasi foreign key (relasi dengan user dan movie)
	User User `gorm:"foreignKey:UserID"`
}