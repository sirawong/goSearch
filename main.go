package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirawong/go-fiber-app/config"
	"github.com/sirawong/go-fiber-app/db"
	"github.com/sirawong/go-fiber-app/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	config.SetFlags()

	db.ConnectDB()

	routes.AppRoutes(app)

	err := app.Listen(":" + config.NewFlags.Port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
