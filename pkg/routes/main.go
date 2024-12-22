package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/tetrex/wecredit-assignment/pkg/services"
)

func InitRoutes(router *echo.Echo, services *services.Services, l *zerolog.Logger) {
	// Public routes (No JWT required)
	router.GET("/", services.Health.HealthCheck)
	// router.POST("/login", services.Auth.Login)
	router.POST("/signup", services.Auth.SignUp)

	// Protected routes (JWT required)
	// protected := router.Group("/v1", jwt.JWTMiddleware)

	// auth
	// protected.POST("/auth/login", services.Auth.Login)
	// protected.POST("/auth/signup", services.Auth.SignUp)

	// Swagger documentation route (no authentication)
	router.GET("docs/*", echoSwagger.WrapHandler)
	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)
}
