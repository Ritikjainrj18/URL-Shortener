package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ritikjainrj18/shorten-url/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Test)
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("err")
	}
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
