package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type Cache struct {
	UUID             string `json:"uuid" gorm:"primaryKey"`
	Type             string `json:"type"`
	ConnectionString string `json:"connection_string"`
	AppUUID          string `json:"app_uuid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *Cache) Map(cache *appService.CacheRequest) {
	c.Type = cache.Type
	c.ConnectionString = cache.ConnectionString
	c.AppUUID = cache.AppUuid
}
