package routes

import (
	"be0/internal/config"
	"be0/internal/handlers"
	"be0/internal/utils/logger"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/labstack/echo/v4"
)

func SetupUploadRoutes(api *echo.Group, cfg *config.Config) {
	log := logger.New("upload_routes")

	// Initialize upload handler
	uploadHandler := handlers.NewUploadHandler(
		types.ObjectCannedACLAuthenticatedRead,
	)

	fileGroup := api.Group("/files")

	fileGroup.POST("/upload", uploadHandler.UploadFile)

	log.Success("Upload routes initialized successfully")
}
