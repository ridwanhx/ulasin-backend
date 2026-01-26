package model

import "time"

// inisialisasi struct
type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(100);not null" json:"nama"`
	Email string `gorm:"varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role string `gorm:"type:varchar(20);default:user" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// inisialisasi foreign key (relasi dengan Reviews)
	Reviews []Review `gorm:"foreignKey:UserID"`
}