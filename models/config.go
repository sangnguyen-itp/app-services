package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type Config struct {
	UUID            string `json:"uuid" gorm:"primaryKey"`
	LogMode         string `json:"log_mode"`
	MigrationDB     string `json:"migration_db"`
	ServerHost      string `json:"server_host"`
	Port            string `json:"port"`
	AccessSecret    string `json:"access_secret"`
	RefreshSecret   string `json:"refresh_secret"`
	AccessDuration  int64  `json:"access_duration"` // second
	RefreshDuration int64  `json:"refresh_duration"`
	AppUUID         string `json:"app_uuid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *Config) Map(config *appService.ConfigRequest) {
	cfg.LogMode = config.LogMode
	cfg.MigrationDB = config.MigrationDb
	cfg.ServerHost = config.ServerHost
	cfg.Port = config.Port
	cfg.AccessSecret = config.AccessSecret
	cfg.RefreshSecret = config.RefreshSecret
	cfg.AccessDuration = config.AccessDuration
	cfg.RefreshDuration = config.RefreshDuration
	cfg.AppUUID = config.AppUuid
}
