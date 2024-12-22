package pkg

import (
	"github.com/rs/zerolog"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	healthService "github.com/tetrex/wecredit-assignment/pkg/health"
	"github.com/tetrex/wecredit-assignment/utils/config"
)

type Services struct {
	Health *healthService.HealthService
}

type ServicesParmas struct {
	config  config.Config
	logger  *zerolog.Logger
	queries *db.Queries
}

func initServices(p ServicesParmas) *Services {
	// init health service
	health_service := healthService.NewHealthService()
	return &Services{
		Health: health_service,
	}
}
