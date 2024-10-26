package routes

import (
	"app/cmd/internal/handlers"
	"app/cmd/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/send-email", handlers.RequestEmailHandler)

	middlewares.EnableCORS(app)

	app.Static("/images", "../public/img")

	middlewares.CSRF(app)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Paginator API",
			"version": "1.0.0"})
	})

	api.Post("/login", handlers.RequestLoginHandler)
	api.Post("/register", handlers.CreateUser)
	api.Post("/forgot-password", handlers.RequestPasswordResetHandler)
	api.Post("/reset-password", handlers.RequestHandleTokenAndNewPasswordHandler)

	middlewares.JWT(app)

	api.Post("/upload-image", handlers.RequestUploadUserImageHandler)
	api.Post("/logout", handlers.RequestLogoutHandler)

	userGroup := api.Group("/user")
	userGroup.Get("/", handlers.RequestUserHandler)

}
