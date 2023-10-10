package routes

import (
	"github.com/catalinfl/godex/handlers"
	"github.com/gofiber/fiber/v2"
)

func SearchRoute(api *fiber.App) {
	router := api.Group("/search")

	router.Get("/:id", handlers.CollyHandler)
}
