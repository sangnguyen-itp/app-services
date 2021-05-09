package implementation

import (
	"context"
	"errors"
	"gRPC/app_services/models"
	"gRPC/app_services/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"github.com/mitchellh/mapstructure"

	appService "gRPC/app_services/internal/services"
)

type IApp interface {
	CreateApp(ctx context.Context, request *appService.AppRequest) (*appService.AppResponse, error)
	GetApp(ctx context.Context, request *appService.GetRequest) (*appService.AppResponse, error)

	CreateConfig(ctx context.Context, request *appService.ConfigRequest) (*appService.ConfigResponse, error)
	GetConfig(ctx context.Context, request *appService.GetRequest) (*appService.ConfigResponse, error)

	CreateDB(ctx context.Context, request *appService.DBInstanceRequest) (*appService.DBInstanceResponse, error)
	GetDB(ctx context.Context, request *appService.GetRequest) (*appService.DBInstanceResponse, error)

	CreateCache(ctx context.Context, request *appService.CacheRequest) (*appService.CacheResponse, error)
	GetCache(ctx context.Context, request *appService.GetRequest) (*appService.CacheResponse, error)

	CreateFirewall(ctx context.Context, request *appService.FirewallRequest) (*appService.FirewallResponse, error)
	GetFirewall(ctx context.Context, request *appService.GetRequest) (*appService.FirewallResponse, error)

	CreatePlatform(ctx context.Context, request *appService.PlatformRequest) (*appService.PlatformResponse, error)
	GetPlatforms(ctx context.Context, request *appService.GetRequest) (*appService.PlatformResponses, error)

	CreateUser(ctx context.Context, request *appService.UserRequest) (*appService.UserResponse, error)
	GetUser(ctx context.Context, request *appService.GetUserRequest) (*appService.UserResponse, error)
}

var _ IApp = &AppImpl{}

type AppImpl struct {
	DB *gorm.DB
}

func (a AppImpl) CreateApp(ctx context.Context, request *appService.AppRequest) (*appService.AppResponse, error) {

	payload := new(models.App)
	payload.Map(request)
	payload.UUID = utils.GenUUID()

	if err := a.DB.Create(payload).Error; err != nil {
		return nil, err
	}

	result := new(appService.AppResponse)
	if err := mapstructure.Decode(payload, &result); err != nil {
		return nil, err
	}
	result.CreatedAt = timestamppb.New(payload.CreatedAt)
	result.UpdatedAt = timestamppb.New(payload.UpdatedAt)

	return result, nil
}

func (a AppImpl) GetApp(ctx context.Context, request *appService.GetRequest) (*appService.AppResponse, error) {
	app := new(models.App)

	if err := a.DB.First(app, "uuid = ?", request.Uuid).Error; err != nil {
		return nil, err
	}

	result := new(appService.AppResponse)
	if err := mapstructure.Decode(app, &result); err != nil {
		return nil, err
	}
	result.CreatedAt = timestamppb.New(app.CreatedAt)
	result.UpdatedAt = timestamppb.New(app.UpdatedAt)

	return result, nil
}

func (a AppImpl) CreateConfig(ctx context.Context, request *appService.ConfigRequest) (*appService.ConfigResponse, error) {
	panic("implement me")
}

func (a AppImpl) GetConfig(ctx context.Context, request *appService.GetRequest) (*appService.ConfigResponse, error) {
	panic("implement me")
}

func (a AppImpl) CreateDB(ctx context.Context, request *appService.DBInstanceRequest) (*appService.DBInstanceResponse, error) {
	panic("implement me")
}

func (a AppImpl) GetDB(ctx context.Context, request *appService.GetRequest) (*appService.DBInstanceResponse, error) {
	panic("implement me")
}

func (a AppImpl) CreateCache(ctx context.Context, request *appService.CacheRequest) (*appService.CacheResponse, error) {
	panic("implement me")
}

func (a AppImpl) GetCache(ctx context.Context, request *appService.GetRequest) (*appService.CacheResponse, error) {
	panic("implement me")
}

func (a AppImpl) CreateFirewall(ctx context.Context, request *appService.FirewallRequest) (*appService.FirewallResponse, error) {
	app := new(models.App)
	if err := a.DB.Preload("Firewall").First(app, "uuid = ?", request.AppUuid).Error; err != nil {
		return nil, errors.New("invalid app_uuid")
	}

	app.Firewall.Map(request)
	if app.Firewall.UUID == "" {
		app.Firewall.UUID = utils.GenUUID()
	}

	if err := a.DB.Save(&app.Firewall).Error; err != nil {
		return nil, err
	}

	result := new(appService.FirewallResponse)
	if err := mapstructure.Decode(app.Firewall, &result); err != nil {
		return nil, err
	}
	result.CreatedAt = timestamppb.New(app.Firewall.CreatedAt)
	result.UpdatedAt = timestamppb.New(app.Firewall.UpdatedAt)

	return result, nil
}

func (a AppImpl) GetFirewall(ctx context.Context, request *appService.GetRequest) (*appService.FirewallResponse, error) {
	app := new(models.App)
	if err := a.DB.Preload("Firewall").First(app, "uuid = ?", request.Uuid).Error; err != nil {
		return nil, err
	}

	result := new(appService.FirewallResponse)
	if err := mapstructure.Decode(app.Firewall, &result); err != nil {
		return nil, err
	}
	result.CreatedAt = timestamppb.New(app.Firewall.CreatedAt)
	result.UpdatedAt = timestamppb.New(app.Firewall.UpdatedAt)

	return result, nil
}

func (a AppImpl) CreatePlatform(ctx context.Context, request *appService.PlatformRequest) (*appService.PlatformResponse, error) {
	panic("implement me")
}

func (a AppImpl) GetPlatforms(ctx context.Context, request *appService.GetRequest) (*appService.PlatformResponses, error) {
	panic("implement me")
}

func (a AppImpl) CreateUser(ctx context.Context, request *appService.UserRequest) (*appService.UserResponse, error) {
	panic("implement me")
}

func (a AppImpl) GetUser(ctx context.Context, request *appService.GetUserRequest) (*appService.UserResponse, error) {
	panic("implement me")
}

func NewAppImpl(db *gorm.DB) *AppImpl {
	return &AppImpl{DB: db}
}
