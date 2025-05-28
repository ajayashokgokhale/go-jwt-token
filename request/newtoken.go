package request

import (
	"jwt-token/actions/newtoken"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewTokenRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	group := router.Group("/newtoken")

	// Register POST /login for customer login.
	group.Post("/", newtoken.HandleNewToken)
}
