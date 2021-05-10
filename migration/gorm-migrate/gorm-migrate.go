package gorm_migrate

import (
	"app-services/models"

	"gorm.io/gorm"
)

func NewMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.App{},
		&models.Firewall{},
		&models.Config{},

		&models.Platform{},
		&models.User{},
	)
}
