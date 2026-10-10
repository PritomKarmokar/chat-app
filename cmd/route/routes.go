package route

import (
	"github.com/PritomKarmokar/chat-app/cmd/middleware"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo) {
	// Apply Correlation ID tracking globally (first middleware for distributed tracing)
	e.Use(middleware.CorrelationID())

	// Apply security headers globally
	e.Use(middleware.SecurityHeaders())

	// Apply CORS for web (configure allowed origins in env)
	e.Use(middleware.CORS())

	// Base prefix for all routes
	basePrefix := e.Group("/chat-app")

	// Service Routes (Health Checks)
	healthGroup := basePrefix.Group("/health")
	RegisterServiceRoutes(healthGroup)
}
