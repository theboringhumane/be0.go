package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Permission scopes
const (
	ScopeAdmin = "admin"
	ScopeRead  = "read"
	ScopeWrite = "create"
)

// ValidateMethodPermission validates if a given scope allows a specific HTTP method
func ValidateMethodPermission(method string, scope string) bool {
	switch scope {
	case ScopeAdmin:
		return true
	case ScopeWrite:
		return method == http.MethodPost || method == http.MethodPut ||
			method == http.MethodDelete || method == http.MethodPatch
	case ScopeRead:
		return method == http.MethodGet
	default:
		return false
	}
}

// GetRequiredPermissionForMethod returns the required permission scope for a given HTTP method
func GetRequiredPermissionForMethod(method string) string {
	switch method {
	case http.MethodGet:
		return ScopeRead
	case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
		return ScopeWrite
	default:
		return ""
	}
}

// RequirePermissions middleware checks if the user/API key has the required permissions
func RequirePermissions(db *gorm.DB, requiredPermissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check if user has admin access first
			if hasAdmin, ok := c.Get("hasAdminAccess").(bool); ok && hasAdmin {
				return next(c)
			}

			method := c.Request().Method

			// For JWT auth, check role-based permissions
			role := c.Get("role").(string)
			scopes := c.Get("scopes").([]string)

			// Admin role has all permissions
			if role == "admin" {
				return next(c)
			}

			// Check if user has any of the required permissions
			hasPermission := false
			for _, scope := range scopes {
				if ValidateMethodPermission(method, scope) {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions")
			}

			return next(c)
		}
	}
}
