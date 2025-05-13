package api

import (
	"be0/internal/api/middleware"
	"be0/internal/api/registry"
	"be0/internal/routes"
	"net/http"

	_ "be0/docs/swagger"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) registerRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Health check
	// @Summary Health check
	// @Description Check if the server is running
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string "OK"
	// @Router /health [get]
	s.echo.GET("/health", s.healthCheck)
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// API v1 group
	api := s.echo.Group("/api/v1")
	auth := middleware.NewAuthMiddleware(s.config.JWT.Secret)
	api.Use(auth.Middleware())

	// Register CRUD routes for all models
	// @Summary Register CRUD routes for all models
	// @Description Register CRUD routes for all models
	registry.RegisterCRUDRoutes(api, s.db)

	routes.SetupUploadRoutes(api, s.config)
}
