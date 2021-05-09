package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type Firewall struct {
	UUID    string `json:"uuid" gorm:"primaryKey"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	AppUUID string `json:"app_uuid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (fw *Firewall) Map(firewall *appService.FirewallRequest) {
	fw.Host = firewall.Host
	fw.Port = firewall.Port
	fw.AppUUID = firewall.AppUuid
}
