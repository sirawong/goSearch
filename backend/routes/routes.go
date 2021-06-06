package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirawong/go-fiber-app/controllers"
)

func AppRoutes(route fiber.Router) {
	route.Post("/api/products/make", controllers.FakerData)
	route.Get("/api/products/get", controllers.GetData)
}
