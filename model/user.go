package model

import "time"

// inisialisasi struct
type User struct {
	ID uint `gorm:"primaryKey`
	Nama string `gorm:"type:varchar(100);not null`
	Email string `gorm:"varchar(100);uniqueIndex;not null`
	Password string `gorm:"type:varchar(255);not null"`
	Role string `gorm:"type:varchar(20);default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// inisialisasi foreign key (relasi dengan Reviews)
	Reviews []Review `gorm:"foreignKey:UserID"`
}