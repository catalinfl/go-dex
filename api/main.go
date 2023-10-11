package main

import (
	"github.com/catalinfl/godex/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))
	routes.SearchRoute(app)

	app.Listen(":3000")
}
