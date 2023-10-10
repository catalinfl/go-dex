package main

import (
	"github.com/catalinfl/godex/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SearchRoute(app)

	app.Listen(":3000")
}
