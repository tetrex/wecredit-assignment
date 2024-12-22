package services

import (
	"github.com/rs/zerolog"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	authService "github.com/tetrex/wecredit-assignment/pkg/auth"
	healthService "github.com/tetrex/wecredit-assignment/pkg/health"
	"github.com/tetrex/wecredit-assignment/utils/config"
)

type Services struct {
	Health *healthService.HealthService
	Auth   *authService.AuthService
}

type ServicesParmas struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func InitServices(p ServicesParmas) *Services {
	// init health service
	health_service := healthService.NewHealthService()
	auth_service := authService.NewAuthService(authService.NewAuthServiceParams{
		Mode:    p.Config.AppEnv,
		Logger:  p.Logger,
		Queries: p.Queries,
	})
	return &Services{
		Health: health_service,
		Auth:   auth_service,
	}
}
