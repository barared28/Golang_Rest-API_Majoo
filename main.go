package main

import (
	"test/server/database"
	"test/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	database.InitDatabase()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":5000")
}
