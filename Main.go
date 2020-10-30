package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	// Load config.yml
	var defaults conf
	defaults.getDefaults()

	callRestEndpoint()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Banana 👋!")
	})

	app.Static("/", "./public")

	// Set default port if not set
	port, isPresent := os.LookupEnv("fiber_port")
	if isPresent == false {

		port = defaults.Port
	}

	app.Listen(":" + port)
}
