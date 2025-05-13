package registry

import (
	"github.com/labstack/echo/v4"

	"be0/internal/api/controllers"
	"be0/internal/api/middleware"
	"be0/internal/models"
	"be0/internal/services"

	"gorm.io/gorm"
)

// 📝 RegisterCRUDRoutes registers CRUD routes for all models - godoc
// @Summary Register CRUD routes for all models
// @Description Register CRUD routes for all models
// @Accept json
// @Produce json
func RegisterCRUDRoutes(g *echo.Group, db *gorm.DB) {
	// Teams
	teamService := services.NewBaseService(db, models.Team{})
	teamController := controllers.NewBaseController(teamService)
	teamGroup := g.Group("/teams")
	teamGroup.Use(middleware.RequirePermissions(db, "teams:read"))

	// @Summary List teams
	// @Description Get a list of all teams
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.Team
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/teams [get]
	teamGroup.GET("", teamController.List)
	// @Summary Get team
	// @Description Get a team by ID
	// @Accept json
	// @Produce json
	// @Param id path string true "Team ID"
	// @Success 200 {object} models.Team
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 404 {object} map[string]string "Not found"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/teams/{id} [get]
	teamGroup.GET("/:id", teamController.Get)

	// Protected team routes
	teamWriteGroup := teamGroup.Group("")
	teamWriteGroup.Use(middleware.RequirePermissions(db, "teams:write"))
	// @Summary Create team
	// @Description Create a new team
	// @Accept json
	// @Produce json
	// @Param team body models.Team true "Team object"
	// @Success 201 {object} models.Team
	// @Failure 400 {object} map[string]string "Bad request"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/teams [post]
	teamWriteGroup.POST("", teamController.Create)
	// @Summary Update team
	// @Description Update an existing team
	// @Accept json
	// @Produce json
	// @Param id path string true "Team ID"
	// @Param team body models.Team true "Team object"
	// @Success 200 {object} models.Team
	// @Failure 400 {object} map[string]string "Bad request"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 404 {object} map[string]string "Not found"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/teams/{id} [put]
	teamWriteGroup.PUT("/:id", teamController.Update)
	// @Summary Delete team
	// @Description Delete a team
	// @Accept json
	// @Produce json
	// @Param id path string true "Team ID"
	// @Success 204 "No content"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 404 {object} map[string]string "Not found"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/teams/{id} [delete]
	teamWriteGroup.DELETE("/:id", teamController.Delete)

	// Team Invitations with team-specific permissions
	invitationService := services.NewBaseService(db, models.TeamInvite{})
	invitationController := controllers.NewBaseController(invitationService)
	invitationGroup := g.Group("/team-invitations")
	invitationGroup.Use(middleware.RequirePermissions(db, "team_invites:read"))
	// @Summary List team invitations
	// @Description Get a list of all team invitations
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.TeamInvite
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/team-invitations [get]
	invitationGroup.GET("", invitationController.List)

	// Protected invitation routes
	invitationWriteGroup := invitationGroup.Group("")
	invitationWriteGroup.Use(middleware.RequirePermissions(db, "team_invites:write"))
	// @Summary Delete team invitation
	// @Description Delete a team invitation
	// @Accept json
	// @Produce json
	// @Param id path string true "Invitation ID"
	// @Success 204 "No content"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 404 {object} map[string]string "Not found"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/team-invitations/{id} [delete]
	invitationWriteGroup.DELETE("/:id", invitationController.Delete)

	// file routes
	fileService := services.NewBaseService(db, models.File{})
	fileController := controllers.NewBaseController(fileService)
	fileGroup := g.Group("/files")
	fileGroup.Use(middleware.RequirePermissions(db, "files:read"))
	// @Summary List files
	// @Description Get a list of all files
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.File
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/files [get]
	fileGroup.GET("", fileController.List)
	// @Summary Get file
	// @Description Get a file by ID
	// @Accept json
	// @Produce json
	// @Param id path string true "File ID"
	// @Success 200 {object} models.File
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Failure 403 {object} map[string]string "Forbidden"
	// @Failure 404 {object} map[string]string "Not found"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/files/{id} [get]
	fileGroup.GET("/:id", fileController.Get)
}
