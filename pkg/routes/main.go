package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/tetrex/wecredit-assignment/pkg/services"
	"github.com/tetrex/wecredit-assignment/utils/jwt"
)

func InitRoutes(router *echo.Echo, services *services.Services, l *zerolog.Logger) {
	// Public routes (No JWT required)
	router.GET("/", services.Health.HealthCheck)

	// auth
	router.POST("/v1/login", services.Auth.Login)
	router.POST("/v1/signup", services.Auth.SignUp)

	// Protected routes (JWT required)
	protected := router.Group("/v1", jwt.JWTMiddleware)

	// auth
	protected.GET("/user", services.User.GetUserById)

	// Swagger documentation route (no authentication)
	router.GET("docs/*", echoSwagger.WrapHandler)
	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)
}
