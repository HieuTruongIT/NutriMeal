package routes

import (
	"nitri-meal-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/signin", handlers.SignIn)
	auth.Post("/signup", handlers.SignUp)
	auth.Post("/signout", handlers.SignOut)

	// User routes
	users := api.Group("/users")
	users.Get("/email", handlers.GetUserByEmail)
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUser)
	users.Put("/:id", handlers.UpdateUser)

	// Health goal routes
	healthGoals := api.Group("/health-goals")
	healthGoals.Post("/", handlers.CreateHealthGoal)
	healthGoals.Get("/user/:userId", handlers.GetHealthGoalsByUserId)
	healthGoals.Get("/:user_id", handlers.GetHealthGoal)
	healthGoals.Put("/:user_id", handlers.UpdateHealthGoal)
}
