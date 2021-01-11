package routes

import (
	"test/server/handlers"
	"test/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes for setup all route
func SetupRoutes(app *fiber.App) {
	// AUTH route
	app.Post("/api/v1/login", handlers.Login)
	app.Post("/api/v1/Register", handlers.Register)

	// User route
	app.Get("api/v1/users", middlewares.Protected(), handlers.GetAllUsers)
	app.Get("api/v1/user/:id", middlewares.Protected(), handlers.GetUserByID)
	app.Put("api/v1/user/:id", middlewares.Protected(), handlers.UpdateUserByID)
	app.Delete("api/v1/user/:id", middlewares.Protected(), handlers.DeleteUserByID)
}
