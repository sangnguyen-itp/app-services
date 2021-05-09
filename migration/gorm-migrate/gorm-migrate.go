package gorm_migrate

import (
	"gRPC/app_services/models"
	"gorm.io/gorm"
)

func NewMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
			&models.Database{},
			&models.Firewall{},
			&models.Cache{},

			&models.Config{},
			&models.App{},

			&models.Platform{},
			&models.User{},
		)
}

