package main

import (
	"debtomate/module/feature/auth"
	"debtomate/module/feature/middleware"
	"debtomate/utils/database"
	"debtomate/utils/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"strconv"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	app.Use(middleware.ConfigureLogging())
	err := viper.ViperConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	db, err := database.InitPGSDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	database.Migrate(db)

	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)

	port := viper.ViperConfig.GetPort()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, DebtoMate")
	})

	err = app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
