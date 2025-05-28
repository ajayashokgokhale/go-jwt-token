package request

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Validate app parameter.
	if app == nil {
		panic("fiber app cannot be nil")
	}

	// Create /api route group for all API endpoints.
	api := app.Group("/api")

	RegisterNewTokenRoutes(api)     // New token generation
	RegisterTokenExtractRoutes(api) // Token extraction

	// Create protected route group with JWT authentication middleware.

	/**
	protected := api.Use(jwtgenx.AuthMiddleware())
	ResetPasswordRoutes(protected) // Password reset (authenticated)
	**/

}
