package main

import (
	"github.com/Levantate-labs/backend-boilerplate/config"
	"github.com/Levantate-labs/backend-boilerplate/routes"
	"github.com/Levantate-labs/backend-boilerplate/storage"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.InitConfig()
	config := config.LoadConfig()
	storage.ConnectDB(config)
}

func main() {
	config := config.LoadConfig()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":" + config.Port)
}
