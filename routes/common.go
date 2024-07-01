package routes

import (
	"github.com/Levantate-labs/backend-boilerplate/controllers"
	"github.com/gofiber/fiber/v2"
)

func CommonRoutes(router fiber.Router) {
	router.Get("/healthcheck", controllers.HealthCheck)
}
