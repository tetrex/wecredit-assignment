package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	db "github.com/tetrex/wecredit-assignment/db/sqlc"
	"github.com/tetrex/wecredit-assignment/pkg/routes"
	"github.com/tetrex/wecredit-assignment/pkg/services"
	"github.com/tetrex/wecredit-assignment/utils/config"
	custommiddleware "github.com/tetrex/wecredit-assignment/utils/custom_middleware"
	"golang.org/x/time/rate"
)

type Server struct {
	config   config.Config
	router   *echo.Echo
	logger   *zerolog.Logger
	queries  *db.Queries
	services *services.Services
}

type ServerParams struct {
	Config    config.Config
	Logger    *zerolog.Logger
	PgQueries *db.Queries
}

func (s *Server) GetServices() *services.Services {
	return s.services
}

func (s *Server) GetConfig() config.Config {
	return s.config
}

func (s *Server) GetRouter() *echo.Echo {
	return s.router
}

func (s *Server) GetLogger() *zerolog.Logger {
	return s.logger
}

func (s *Server) GetQueries() *db.Queries {
	return s.queries
}

func NewServer(c *ServerParams) (*Server, error) {
	router := echo.New()

	// for unique reqiest id
	router.Use(middleware.RequestID())

	// device id
	router.Use(custommiddleware.DeviceIDMiddleware())

	// logger
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, request_id=${id}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency_nano=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out}\n",
	}))

	// stack trace for debugging
	router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	// rate limmiter
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(100), Burst: 30, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	router.Use(middleware.RateLimiterWithConfig(config))

	// cors
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000/", "*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Content-Type", "Authorization", "Accept", "Origin", "X-Requested-With",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Accept-Language", "Content-Language",
			"Cookie"},
		Skipper: func(c echo.Context) bool {
			// adding this so we can skip postamm request from blocking
			h := c.Request().Header.Get("User-Agent")
			user_agent := strings.Split(h, "/")

			return strings.EqualFold("PostmanRuntime", user_agent[0])

		},
		ExposeHeaders: []string{"Set-Cookie"},
	}))
	// services
	services := services.InitServices(services.ServicesParmas{
		Config:  c.Config,
		Logger:  c.Logger,
		Queries: c.PgQueries,
	})

	// routes setup
	routes.InitRoutes(router, services, c.Logger)

	return &Server{
		config:   c.Config,
		router:   router,
		logger:   c.Logger,
		queries:  c.PgQueries,
		services: services,
	}, nil
}
