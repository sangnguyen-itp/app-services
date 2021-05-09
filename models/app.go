package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type App struct {
	UUID          string `json:"uuid" gorm:"primaryKey"`
	Name          string `json:"name"`
	LatestVersion string `json:"latest_version"`
	RunningStatus string `json:"running_status"`
	Type          string `json:"type"`

	Platforms []*Platform `json:"platforms" gorm:"foreignKey:AppUUID"`
	Firewall  Firewall    `json:"firewall" gorm:"foreignKey:AppUUID"`
	Database  Database    `json:"database" gorm:"foreignKey:AppUUID"`
	Cache     Cache       `json:"cache" gorm:"foreignKey:AppUUID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (app *App) Map(appRequest *appService.AppRequest) {
	app.Name = appRequest.Name
	app.LatestVersion = appRequest.LatestVersion
	app.RunningStatus = appRequest.RunningStatus
	app.Type = appRequest.Type
}
