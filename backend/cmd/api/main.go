package main

import (
	"app/cmd/internal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	routes.Routes(app)

	fmt.Println("Starting server on :3000")

	app.Listen(":3000")
}
