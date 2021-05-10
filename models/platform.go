package models

import (
	appService "app-services/internal/services"
	"time"
)

type Platform struct {
	UUID   string `json:"uuid" gorm:"primaryKey"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`

	AppUUID string `json:"app_uuid"`
	App     App    `json:"app" gorm:"foreignKey:AppUUID"`

	Users []User `json:"users" gorm:"many2many:user_platforms"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (pf *Platform) Map(platform *appService.PlatformRequest) {
	pf.AppUUID = platform.AppUuid
	pf.Name = platform.Name
	pf.Code = platform.Code
	pf.Status = platform.Status
}
