package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type Database struct {
	UUID             string `json:"uuid" gorm:"primaryKey"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	ConnectionString string `json:"connection_string"`
	AppUUID          string `json:"app_uuid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (db *Database) Map(dbInstance *appService.DBInstanceRequest) {
	db.Name = dbInstance.Name
	db.Type = dbInstance.Type
	db.ConnectionString = dbInstance.ConnectionString
	db.AppUUID = dbInstance.AppUuid
}
