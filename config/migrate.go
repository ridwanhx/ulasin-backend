package config

import "ulasin-backend/model"

// auto migration
func Migrate() error {
	// migrasi berdasarkan urutan
	return DB.AutoMigrate(
		&model.User{},
		&model.Movie{},
		&model.Review{},
	)
}